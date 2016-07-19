// This file was generated by counterfeiter
package plugininstallerfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/actors/plugininstaller"
)

type FakePluginInstaller struct {
	InstallStub        func(inputSourceFilepath string) string
	installMutex       sync.RWMutex
	installArgsForCall []struct {
		inputSourceFilepath string
	}
	installReturns struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePluginInstaller) Install(inputSourceFilepath string) string {
	fake.installMutex.Lock()
	fake.installArgsForCall = append(fake.installArgsForCall, struct {
		inputSourceFilepath string
	}{inputSourceFilepath})
	fake.recordInvocation("Install", []interface{}{inputSourceFilepath})
	fake.installMutex.Unlock()
	if fake.InstallStub != nil {
		return fake.InstallStub(inputSourceFilepath)
	} else {
		return fake.installReturns.result1
	}
}

func (fake *FakePluginInstaller) InstallCallCount() int {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return len(fake.installArgsForCall)
}

func (fake *FakePluginInstaller) InstallArgsForCall(i int) string {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return fake.installArgsForCall[i].inputSourceFilepath
}

func (fake *FakePluginInstaller) InstallReturns(result1 string) {
	fake.InstallStub = nil
	fake.installReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePluginInstaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePluginInstaller) recordInvocation(key string, args []interface{}) {
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

var _ plugininstaller.PluginInstaller = new(FakePluginInstaller)
