// Copyright (C) 2018-Present Pivotal Software, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under the terms of the under the Apache License,
// Version 2.0 (the "License”); you may not use this file except in compliance with the License. You may obtain a copy
// of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package find_bindings

import (
	"net/url"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

//go:generate counterfeiter . CFClient
type CFClient interface {
	GetAppByGuid(guid string) (cfclient.App, error)
	GetOrgByGuid(spaceGUID string) (cfclient.Org, error)
	GetSpaceByGuid(spaceGUID string) (cfclient.Space, error)
	ListServicesByQuery(url.Values) ([]cfclient.Service, error)
	ListServiceBindingsByQuery(query url.Values) ([]cfclient.ServiceBinding, error)
	ListServicePlansByQuery(query url.Values) ([]cfclient.ServicePlan, error)
	ListServiceKeysByQuery(query url.Values) ([]cfclient.ServiceKey, error)
	ListServiceInstancesByQuery(query url.Values) ([]cfclient.ServiceInstance, error)
}

type Binding struct {
	Name                string
	ServiceInstanceName string
	ServiceInstanceGuid string
	OrgName             string
	SpaceName           string
	Type                string
}

//go:generate counterfeiter . BindingFinder
type BindingFinder interface {
	FindBindings(serviceLabel string) ([]Binding, error)
}

type bindingFinder struct {
	cfClient CFClient
}

func NewBindingFinder(cfClient CFClient) BindingFinder {
	return &bindingFinder{
		cfClient: cfClient,
	}
}

func (bf *bindingFinder) FindBindings(serviceLabel string) ([]Binding, error) {
	serviceGUID, err := bf.serviceGUIDForLabel(serviceLabel)
	if err != nil {
		return nil, errors.Wrapf(err, `failed to lookup service matching label %q`, serviceLabel)
	}

	servicePlans, err := bf.servicePlansForServiceGUID(serviceGUID)
	if err != nil {
		return nil, errors.Wrapf(err, `failed to lookup service plans for service (guid: %q, label: %q)`, serviceGUID, serviceLabel)
	}

	var (
		result []Binding
		errs   error
	)

	serviceInstances, err := bf.serviceInstancesForServicePlans(servicePlans)
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	for _, instance := range serviceInstances {
		bindings, err := bf.listServiceBindingsForInstance(instance)
		if err != nil {
			errs = multierror.Append(errs, err)
		}
		result = append(result, bindings...)

		bindings, err = bf.listServiceKeysForInstance(instance)
		if err != nil {
			errs = multierror.Append(errs, err)
		}
		result = append(result, bindings...)
	}

	return result, errs
}

func (bf *bindingFinder) serviceGUIDForLabel(serviceLabel string) (serviceGUID string, err error) {
	query := url.Values{}
	query.Set("q", "label:"+serviceLabel)
	services, err := bf.cfClient.ListServicesByQuery(query)
	if err != nil {
		return "", err
	}

	switch len(services) {
	case 0:
		return "", errors.New("no matching services found")
	case 1:
		return services[0].Guid, nil
	default:
		return "", errors.Errorf("found %d matching services, expected 1", len(services))
	}
}

func (bf *bindingFinder) servicePlansForServiceGUID(serviceGUID string) (servicePlans []cfclient.ServicePlan, err error) {
	query := url.Values{}
	query.Set("q", "service_guid:"+serviceGUID)
	return bf.cfClient.ListServicePlansByQuery(query)
}

func (bf *bindingFinder) serviceInstancesForServicePlans(servicePlans []cfclient.ServicePlan) ([]cfclient.ServiceInstance, error) {
	var (
		result []cfclient.ServiceInstance
		errs   error
	)

	for _, plan := range servicePlans {
		instances, err := bf.cfClient.ListServiceInstancesByQuery(url.Values{
			"q": []string{
				"service_plan_guid:" + plan.Guid,
			},
		})
		if err != nil {
			errs = multierror.Append(
				errs,
				errors.Wrapf(err, `failed to lookup service instances for service plan (name: %q, guid: %q)`, plan.Name, plan.Guid),
			)
		}

		result = append(result, instances...)
	}

	return result, errs
}

func (bf *bindingFinder) listServiceBindingsForInstance(instance cfclient.ServiceInstance) ([]Binding, error) {
	query := url.Values{}
	query.Set("q", "service_instance_guid:"+instance.Guid)
	serviceBindings, err := bf.cfClient.ListServiceBindingsByQuery(query)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve service bindings for service instance (name: %q guid: %q)", instance.Name, instance.Guid)
	}

	var (
		result []Binding
		errs   error
	)

	for _, b := range serviceBindings {
		if err != nil {
			errs = multierror.Append(errs, errors.Wrapf(err, "failed to resolve secure binding credentials for app (guid: %q)", b.AppGuid))
			continue
		}

		app, err := bf.cfClient.GetAppByGuid(b.AppGuid)
		if err != nil {
			errs = multierror.Append(errs, errors.Wrapf(err, "failed to lookup app info for app guid %q", b.AppGuid))
			continue
		}

		result = append(result, Binding{
			Name:                app.Name,
			ServiceInstanceName: instance.Name,
			ServiceInstanceGuid: instance.Guid,
			OrgName:             app.SpaceData.Entity.OrgData.Entity.Name,
			SpaceName:           app.SpaceData.Entity.Name,
			Type:                "AppBinding",
		})
	}

	return result, errs
}

func (bf *bindingFinder) listServiceKeysForInstance(instance cfclient.ServiceInstance) ([]Binding, error) {
	query := url.Values{}
	query.Set("q", "service_instance_guid:"+instance.Guid)
	serviceKeys, err := bf.cfClient.ListServiceKeysByQuery(query)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve service keys for service instance (name: %q guid: %q)", instance.Name, instance.Guid)
	}

	var (
		result []Binding
		errs   error
	)

	for _, k := range serviceKeys {
		space, err := bf.spaceDataForGUID(instance.SpaceGuid)
		if err != nil {
			// XXX: Think about the error and what information to surface here
			errs = multierror.Append(errs, errors.Wrapf(err, "failed to lookup space info for service key (name: %q instance-guid: %q)", k.Name, instance.Guid))
			continue
		}

		result = append(result, Binding{
			Name:                k.Name,
			ServiceInstanceName: instance.Name,
			ServiceInstanceGuid: instance.Guid,
			OrgName:             space.OrgData.Entity.Name,
			SpaceName:           space.Name,
			Type:                "ServiceKeyBinding",
		})
	}

	return result, errs
}

func (bf *bindingFinder) spaceDataForGUID(spaceGUID string) (cfclient.Space, error) {
	space, err := bf.cfClient.GetSpaceByGuid(spaceGUID)
	if err != nil {
		return space, err
	}

	org, err := bf.cfClient.GetOrgByGuid(space.OrganizationGuid)
	if err != nil {
		return space, err
	}

	space.OrgData.Entity = org

	return space, nil
}
