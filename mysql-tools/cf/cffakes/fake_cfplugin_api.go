// Code generated by counterfeiter. DO NOT EDIT.
package cffakes

import (
	"sync"

	plugin_models "code.cloudfoundry.org/cli/plugin/models"
	"github.com/pivotal-cf/mysql-cli-plugin/mysql-tools/cf"
)

type FakeCFPluginAPI struct {
	AccessTokenStub        func() (string, error)
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct {
	}
	accessTokenReturns struct {
		result1 string
		result2 error
	}
	accessTokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CliCommandStub        func(...string) ([]string, error)
	cliCommandMutex       sync.RWMutex
	cliCommandArgsForCall []struct {
		arg1 []string
	}
	cliCommandReturns struct {
		result1 []string
		result2 error
	}
	cliCommandReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	CliCommandWithoutTerminalOutputStub        func(...string) ([]string, error)
	cliCommandWithoutTerminalOutputMutex       sync.RWMutex
	cliCommandWithoutTerminalOutputArgsForCall []struct {
		arg1 []string
	}
	cliCommandWithoutTerminalOutputReturns struct {
		result1 []string
		result2 error
	}
	cliCommandWithoutTerminalOutputReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetCurrentSpaceStub        func() (plugin_models.Space, error)
	getCurrentSpaceMutex       sync.RWMutex
	getCurrentSpaceArgsForCall []struct {
	}
	getCurrentSpaceReturns struct {
		result1 plugin_models.Space
		result2 error
	}
	getCurrentSpaceReturnsOnCall map[int]struct {
		result1 plugin_models.Space
		result2 error
	}
	GetServiceStub        func(string) (plugin_models.GetService_Model, error)
	getServiceMutex       sync.RWMutex
	getServiceArgsForCall []struct {
		arg1 string
	}
	getServiceReturns struct {
		result1 plugin_models.GetService_Model
		result2 error
	}
	getServiceReturnsOnCall map[int]struct {
		result1 plugin_models.GetService_Model
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFPluginAPI) AccessToken() (string, error) {
	fake.accessTokenMutex.Lock()
	ret, specificReturn := fake.accessTokenReturnsOnCall[len(fake.accessTokenArgsForCall)]
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct {
	}{})
	stub := fake.AccessTokenStub
	fakeReturns := fake.accessTokenReturns
	fake.recordInvocation("AccessToken", []interface{}{})
	fake.accessTokenMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFPluginAPI) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeCFPluginAPI) AccessTokenCalls(stub func() (string, error)) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = stub
}

func (fake *FakeCFPluginAPI) AccessTokenReturns(result1 string, result2 error) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCFPluginAPI) AccessTokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = nil
	if fake.accessTokenReturnsOnCall == nil {
		fake.accessTokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.accessTokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCFPluginAPI) CliCommand(arg1 ...string) ([]string, error) {
	fake.cliCommandMutex.Lock()
	ret, specificReturn := fake.cliCommandReturnsOnCall[len(fake.cliCommandArgsForCall)]
	fake.cliCommandArgsForCall = append(fake.cliCommandArgsForCall, struct {
		arg1 []string
	}{arg1})
	stub := fake.CliCommandStub
	fakeReturns := fake.cliCommandReturns
	fake.recordInvocation("CliCommand", []interface{}{arg1})
	fake.cliCommandMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFPluginAPI) CliCommandCallCount() int {
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	return len(fake.cliCommandArgsForCall)
}

func (fake *FakeCFPluginAPI) CliCommandCalls(stub func(...string) ([]string, error)) {
	fake.cliCommandMutex.Lock()
	defer fake.cliCommandMutex.Unlock()
	fake.CliCommandStub = stub
}

func (fake *FakeCFPluginAPI) CliCommandArgsForCall(i int) []string {
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	argsForCall := fake.cliCommandArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFPluginAPI) CliCommandReturns(result1 []string, result2 error) {
	fake.cliCommandMutex.Lock()
	defer fake.cliCommandMutex.Unlock()
	fake.CliCommandStub = nil
	fake.cliCommandReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCFPluginAPI) CliCommandReturnsOnCall(i int, result1 []string, result2 error) {
	fake.cliCommandMutex.Lock()
	defer fake.cliCommandMutex.Unlock()
	fake.CliCommandStub = nil
	if fake.cliCommandReturnsOnCall == nil {
		fake.cliCommandReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.cliCommandReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCFPluginAPI) CliCommandWithoutTerminalOutput(arg1 ...string) ([]string, error) {
	fake.cliCommandWithoutTerminalOutputMutex.Lock()
	ret, specificReturn := fake.cliCommandWithoutTerminalOutputReturnsOnCall[len(fake.cliCommandWithoutTerminalOutputArgsForCall)]
	fake.cliCommandWithoutTerminalOutputArgsForCall = append(fake.cliCommandWithoutTerminalOutputArgsForCall, struct {
		arg1 []string
	}{arg1})
	stub := fake.CliCommandWithoutTerminalOutputStub
	fakeReturns := fake.cliCommandWithoutTerminalOutputReturns
	fake.recordInvocation("CliCommandWithoutTerminalOutput", []interface{}{arg1})
	fake.cliCommandWithoutTerminalOutputMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFPluginAPI) CliCommandWithoutTerminalOutputCallCount() int {
	fake.cliCommandWithoutTerminalOutputMutex.RLock()
	defer fake.cliCommandWithoutTerminalOutputMutex.RUnlock()
	return len(fake.cliCommandWithoutTerminalOutputArgsForCall)
}

func (fake *FakeCFPluginAPI) CliCommandWithoutTerminalOutputCalls(stub func(...string) ([]string, error)) {
	fake.cliCommandWithoutTerminalOutputMutex.Lock()
	defer fake.cliCommandWithoutTerminalOutputMutex.Unlock()
	fake.CliCommandWithoutTerminalOutputStub = stub
}

func (fake *FakeCFPluginAPI) CliCommandWithoutTerminalOutputArgsForCall(i int) []string {
	fake.cliCommandWithoutTerminalOutputMutex.RLock()
	defer fake.cliCommandWithoutTerminalOutputMutex.RUnlock()
	argsForCall := fake.cliCommandWithoutTerminalOutputArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFPluginAPI) CliCommandWithoutTerminalOutputReturns(result1 []string, result2 error) {
	fake.cliCommandWithoutTerminalOutputMutex.Lock()
	defer fake.cliCommandWithoutTerminalOutputMutex.Unlock()
	fake.CliCommandWithoutTerminalOutputStub = nil
	fake.cliCommandWithoutTerminalOutputReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCFPluginAPI) CliCommandWithoutTerminalOutputReturnsOnCall(i int, result1 []string, result2 error) {
	fake.cliCommandWithoutTerminalOutputMutex.Lock()
	defer fake.cliCommandWithoutTerminalOutputMutex.Unlock()
	fake.CliCommandWithoutTerminalOutputStub = nil
	if fake.cliCommandWithoutTerminalOutputReturnsOnCall == nil {
		fake.cliCommandWithoutTerminalOutputReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.cliCommandWithoutTerminalOutputReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCFPluginAPI) GetCurrentSpace() (plugin_models.Space, error) {
	fake.getCurrentSpaceMutex.Lock()
	ret, specificReturn := fake.getCurrentSpaceReturnsOnCall[len(fake.getCurrentSpaceArgsForCall)]
	fake.getCurrentSpaceArgsForCall = append(fake.getCurrentSpaceArgsForCall, struct {
	}{})
	stub := fake.GetCurrentSpaceStub
	fakeReturns := fake.getCurrentSpaceReturns
	fake.recordInvocation("GetCurrentSpace", []interface{}{})
	fake.getCurrentSpaceMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFPluginAPI) GetCurrentSpaceCallCount() int {
	fake.getCurrentSpaceMutex.RLock()
	defer fake.getCurrentSpaceMutex.RUnlock()
	return len(fake.getCurrentSpaceArgsForCall)
}

func (fake *FakeCFPluginAPI) GetCurrentSpaceCalls(stub func() (plugin_models.Space, error)) {
	fake.getCurrentSpaceMutex.Lock()
	defer fake.getCurrentSpaceMutex.Unlock()
	fake.GetCurrentSpaceStub = stub
}

func (fake *FakeCFPluginAPI) GetCurrentSpaceReturns(result1 plugin_models.Space, result2 error) {
	fake.getCurrentSpaceMutex.Lock()
	defer fake.getCurrentSpaceMutex.Unlock()
	fake.GetCurrentSpaceStub = nil
	fake.getCurrentSpaceReturns = struct {
		result1 plugin_models.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFPluginAPI) GetCurrentSpaceReturnsOnCall(i int, result1 plugin_models.Space, result2 error) {
	fake.getCurrentSpaceMutex.Lock()
	defer fake.getCurrentSpaceMutex.Unlock()
	fake.GetCurrentSpaceStub = nil
	if fake.getCurrentSpaceReturnsOnCall == nil {
		fake.getCurrentSpaceReturnsOnCall = make(map[int]struct {
			result1 plugin_models.Space
			result2 error
		})
	}
	fake.getCurrentSpaceReturnsOnCall[i] = struct {
		result1 plugin_models.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFPluginAPI) GetService(arg1 string) (plugin_models.GetService_Model, error) {
	fake.getServiceMutex.Lock()
	ret, specificReturn := fake.getServiceReturnsOnCall[len(fake.getServiceArgsForCall)]
	fake.getServiceArgsForCall = append(fake.getServiceArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetServiceStub
	fakeReturns := fake.getServiceReturns
	fake.recordInvocation("GetService", []interface{}{arg1})
	fake.getServiceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFPluginAPI) GetServiceCallCount() int {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	return len(fake.getServiceArgsForCall)
}

func (fake *FakeCFPluginAPI) GetServiceCalls(stub func(string) (plugin_models.GetService_Model, error)) {
	fake.getServiceMutex.Lock()
	defer fake.getServiceMutex.Unlock()
	fake.GetServiceStub = stub
}

func (fake *FakeCFPluginAPI) GetServiceArgsForCall(i int) string {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	argsForCall := fake.getServiceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCFPluginAPI) GetServiceReturns(result1 plugin_models.GetService_Model, result2 error) {
	fake.getServiceMutex.Lock()
	defer fake.getServiceMutex.Unlock()
	fake.GetServiceStub = nil
	fake.getServiceReturns = struct {
		result1 plugin_models.GetService_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCFPluginAPI) GetServiceReturnsOnCall(i int, result1 plugin_models.GetService_Model, result2 error) {
	fake.getServiceMutex.Lock()
	defer fake.getServiceMutex.Unlock()
	fake.GetServiceStub = nil
	if fake.getServiceReturnsOnCall == nil {
		fake.getServiceReturnsOnCall = make(map[int]struct {
			result1 plugin_models.GetService_Model
			result2 error
		})
	}
	fake.getServiceReturnsOnCall[i] = struct {
		result1 plugin_models.GetService_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCFPluginAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	fake.cliCommandWithoutTerminalOutputMutex.RLock()
	defer fake.cliCommandWithoutTerminalOutputMutex.RUnlock()
	fake.getCurrentSpaceMutex.RLock()
	defer fake.getCurrentSpaceMutex.RUnlock()
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFPluginAPI) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cf.CFPluginAPI = new(FakeCFPluginAPI)
