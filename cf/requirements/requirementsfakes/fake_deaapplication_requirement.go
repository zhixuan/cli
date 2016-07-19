// This file was generated by counterfeiter
package requirementsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/models"
	"code.cloudfoundry.org/cli/cf/requirements"
)

type FakeDEAApplicationRequirement struct {
	ExecuteStub        func() error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct{}
	executeReturns     struct {
		result1 error
	}
	GetApplicationStub        func() models.Application
	getApplicationMutex       sync.RWMutex
	getApplicationArgsForCall []struct{}
	getApplicationReturns     struct {
		result1 models.Application
	}
}

func (fake *FakeDEAApplicationRequirement) Execute() error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeDEAApplicationRequirement) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeDEAApplicationRequirement) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDEAApplicationRequirement) GetApplication() models.Application {
	fake.getApplicationMutex.Lock()
	fake.getApplicationArgsForCall = append(fake.getApplicationArgsForCall, struct{}{})
	fake.getApplicationMutex.Unlock()
	if fake.GetApplicationStub != nil {
		return fake.GetApplicationStub()
	} else {
		return fake.getApplicationReturns.result1
	}
}

func (fake *FakeDEAApplicationRequirement) GetApplicationCallCount() int {
	fake.getApplicationMutex.RLock()
	defer fake.getApplicationMutex.RUnlock()
	return len(fake.getApplicationArgsForCall)
}

func (fake *FakeDEAApplicationRequirement) GetApplicationReturns(result1 models.Application) {
	fake.GetApplicationStub = nil
	fake.getApplicationReturns = struct {
		result1 models.Application
	}{result1}
}

var _ requirements.DEAApplicationRequirement = new(FakeDEAApplicationRequirement)
