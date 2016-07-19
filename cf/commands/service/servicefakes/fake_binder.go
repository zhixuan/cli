// This file was generated by counterfeiter
package servicefakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/commands/service"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeBinder struct {
	BindApplicationStub        func(app models.Application, serviceInstance models.ServiceInstance, paramsMap map[string]interface{}) (apiErr error)
	bindApplicationMutex       sync.RWMutex
	bindApplicationArgsForCall []struct {
		app             models.Application
		serviceInstance models.ServiceInstance
		paramsMap       map[string]interface{}
	}
	bindApplicationReturns struct {
		result1 error
	}
}

func (fake *FakeBinder) BindApplication(app models.Application, serviceInstance models.ServiceInstance, paramsMap map[string]interface{}) (apiErr error) {
	fake.bindApplicationMutex.Lock()
	fake.bindApplicationArgsForCall = append(fake.bindApplicationArgsForCall, struct {
		app             models.Application
		serviceInstance models.ServiceInstance
		paramsMap       map[string]interface{}
	}{app, serviceInstance, paramsMap})
	fake.bindApplicationMutex.Unlock()
	if fake.BindApplicationStub != nil {
		return fake.BindApplicationStub(app, serviceInstance, paramsMap)
	} else {
		return fake.bindApplicationReturns.result1
	}
}

func (fake *FakeBinder) BindApplicationCallCount() int {
	fake.bindApplicationMutex.RLock()
	defer fake.bindApplicationMutex.RUnlock()
	return len(fake.bindApplicationArgsForCall)
}

func (fake *FakeBinder) BindApplicationArgsForCall(i int) (models.Application, models.ServiceInstance, map[string]interface{}) {
	fake.bindApplicationMutex.RLock()
	defer fake.bindApplicationMutex.RUnlock()
	return fake.bindApplicationArgsForCall[i].app, fake.bindApplicationArgsForCall[i].serviceInstance, fake.bindApplicationArgsForCall[i].paramsMap
}

func (fake *FakeBinder) BindApplicationReturns(result1 error) {
	fake.BindApplicationStub = nil
	fake.bindApplicationReturns = struct {
		result1 error
	}{result1}
}

var _ service.Binder = new(FakeBinder)
