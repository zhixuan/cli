// This file was generated by counterfeiter
package repositoryfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/v3/models"
	"code.cloudfoundry.org/cli/cf/v3/repository"
)

type FakeRepository struct {
	GetApplicationsStub        func() ([]models.V3Application, error)
	getApplicationsMutex       sync.RWMutex
	getApplicationsArgsForCall []struct{}
	getApplicationsReturns     struct {
		result1 []models.V3Application
		result2 error
	}
	GetProcessesStub        func(path string) ([]models.V3Process, error)
	getProcessesMutex       sync.RWMutex
	getProcessesArgsForCall []struct {
		path string
	}
	getProcessesReturns struct {
		result1 []models.V3Process
		result2 error
	}
	GetRoutesStub        func(path string) ([]models.V3Route, error)
	getRoutesMutex       sync.RWMutex
	getRoutesArgsForCall []struct {
		path string
	}
	getRoutesReturns struct {
		result1 []models.V3Route
		result2 error
	}
}

func (fake *FakeRepository) GetApplications() ([]models.V3Application, error) {
	fake.getApplicationsMutex.Lock()
	fake.getApplicationsArgsForCall = append(fake.getApplicationsArgsForCall, struct{}{})
	fake.getApplicationsMutex.Unlock()
	if fake.GetApplicationsStub != nil {
		return fake.GetApplicationsStub()
	} else {
		return fake.getApplicationsReturns.result1, fake.getApplicationsReturns.result2
	}
}

func (fake *FakeRepository) GetApplicationsCallCount() int {
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	return len(fake.getApplicationsArgsForCall)
}

func (fake *FakeRepository) GetApplicationsReturns(result1 []models.V3Application, result2 error) {
	fake.GetApplicationsStub = nil
	fake.getApplicationsReturns = struct {
		result1 []models.V3Application
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetProcesses(path string) ([]models.V3Process, error) {
	fake.getProcessesMutex.Lock()
	fake.getProcessesArgsForCall = append(fake.getProcessesArgsForCall, struct {
		path string
	}{path})
	fake.getProcessesMutex.Unlock()
	if fake.GetProcessesStub != nil {
		return fake.GetProcessesStub(path)
	} else {
		return fake.getProcessesReturns.result1, fake.getProcessesReturns.result2
	}
}

func (fake *FakeRepository) GetProcessesCallCount() int {
	fake.getProcessesMutex.RLock()
	defer fake.getProcessesMutex.RUnlock()
	return len(fake.getProcessesArgsForCall)
}

func (fake *FakeRepository) GetProcessesArgsForCall(i int) string {
	fake.getProcessesMutex.RLock()
	defer fake.getProcessesMutex.RUnlock()
	return fake.getProcessesArgsForCall[i].path
}

func (fake *FakeRepository) GetProcessesReturns(result1 []models.V3Process, result2 error) {
	fake.GetProcessesStub = nil
	fake.getProcessesReturns = struct {
		result1 []models.V3Process
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetRoutes(path string) ([]models.V3Route, error) {
	fake.getRoutesMutex.Lock()
	fake.getRoutesArgsForCall = append(fake.getRoutesArgsForCall, struct {
		path string
	}{path})
	fake.getRoutesMutex.Unlock()
	if fake.GetRoutesStub != nil {
		return fake.GetRoutesStub(path)
	} else {
		return fake.getRoutesReturns.result1, fake.getRoutesReturns.result2
	}
}

func (fake *FakeRepository) GetRoutesCallCount() int {
	fake.getRoutesMutex.RLock()
	defer fake.getRoutesMutex.RUnlock()
	return len(fake.getRoutesArgsForCall)
}

func (fake *FakeRepository) GetRoutesArgsForCall(i int) string {
	fake.getRoutesMutex.RLock()
	defer fake.getRoutesMutex.RUnlock()
	return fake.getRoutesArgsForCall[i].path
}

func (fake *FakeRepository) GetRoutesReturns(result1 []models.V3Route, result2 error) {
	fake.GetRoutesStub = nil
	fake.getRoutesReturns = struct {
		result1 []models.V3Route
		result2 error
	}{result1, result2}
}

var _ repository.Repository = new(FakeRepository)
