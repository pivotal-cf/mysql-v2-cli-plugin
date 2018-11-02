package smoke_tests_test

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf/mysql-cli-plugin/test_helpers"
)

var _ = Describe("SmokeTests", func() {
	Context("When a valid donor service instance exists", func() {
		var (
			appDomain          string
			destPlan           string
			springAppName      string
			instanceName       string
			sourceInstanceGUID string
		)

		BeforeEach(func() {
			destPlan = os.Getenv("RECIPIENT_PLAN_NAME")
			appDomain = Config.AppsDomain

			instanceName = generator.PrefixedRandomName("MYSQL", "MIGRATE_SOURCE")
			test_helpers.CreateService(
				os.Getenv("DONOR_SERVICE_NAME"),
				os.Getenv("DONOR_PLAN_NAME"),
				instanceName,
			)
			test_helpers.WaitForService(instanceName, `[Ss]tatus:\s+create succeeded`)
			sourceInstanceGUID = test_helpers.InstanceUUID(instanceName)
		})

		AfterEach(func() {
			if springAppName != "" {
				test_helpers.DeleteApp(springAppName)
			}

			test_helpers.DeleteService(instanceName + "-old")
			test_helpers.DeleteService(instanceName + "-new")
			test_helpers.DeleteService(instanceName)
			test_helpers.WaitForService(
				instanceName+"-old",
				fmt.Sprintf("Service instance %s-old not found", instanceName))
			test_helpers.WaitForService(
				instanceName+"-new",
				fmt.Sprintf("Service instance %s-new not found", instanceName))
			test_helpers.WaitForService(
				instanceName,
				fmt.Sprintf("Service instance %s not found", instanceName))
		})

		It("migrates data from donor to recipient", func() {
			var (
				readValue    string
				springAppURI string
				albumID      string
				writeValue   string
			)

			By("Writing data to the source instance", func() {
				springAppName = generator.PrefixedRandomName("MYSQL", "APP")
				test_helpers.PushApp(springAppName, "../assets/spring-music")

				test_helpers.BindAppToService(springAppName, instanceName)
				test_helpers.StartApp(springAppName)

				springAppURI = springAppName + "." + appDomain
				test_helpers.CheckAppInfo(true, springAppURI, instanceName)

				writeValue = "DM Greatest Hits"
				albumID = test_helpers.WriteData(true, springAppURI, writeValue)
				readValue = test_helpers.ReadData(true, springAppURI, albumID)

				Expect(readValue).To(Equal(writeValue))

				test_helpers.UnbindAppFromService(springAppName, instanceName)
			})

			By("Migrating data using the migrate command", func() {
				cmd := exec.Command("cf", "mysql-tools", "migrate", instanceName, destPlan)
				session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session, "20m", "1s").Should(gexec.Exit(0))
			})

			By("Verifying the destination service was renamed to the source's name", func() {
				Expect(sourceInstanceGUID).NotTo(Equal(test_helpers.InstanceUUID(instanceName)))
			})

			By("Binding the app to the newly created destination instance and reading back data", func() {
				test_helpers.BindAppToService(springAppName, instanceName)
				test_helpers.ExecuteCfCmd("restage", springAppName)

				readValue = test_helpers.ReadData(true, springAppURI, albumID)
				Expect(readValue).To(Equal(writeValue))
			})
		})
	})
})