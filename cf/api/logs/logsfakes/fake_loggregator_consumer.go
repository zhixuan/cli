// This file was generated by counterfeiter
package logsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api/logs"
	"github.com/cloudfoundry/loggregator_consumer"
	"github.com/cloudfoundry/loggregatorlib/logmessage"
)

type FakeLoggregatorConsumer struct {
	TailStub        func(appGUID string, authToken string) (<-chan *logmessage.LogMessage, error)
	tailMutex       sync.RWMutex
	tailArgsForCall []struct {
		appGUID   string
		authToken string
	}
	tailReturns struct {
		result1 <-chan *logmessage.LogMessage
		result2 error
	}
	RecentStub        func(appGUID string, authToken string) ([]*logmessage.LogMessage, error)
	recentMutex       sync.RWMutex
	recentArgsForCall []struct {
		appGUID   string
		authToken string
	}
	recentReturns struct {
		result1 []*logmessage.LogMessage
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	SetOnConnectCallbackStub        func(func())
	setOnConnectCallbackMutex       sync.RWMutex
	setOnConnectCallbackArgsForCall []struct {
		arg1 func()
	}
	SetDebugPrinterStub        func(loggregator_consumer.DebugPrinter)
	setDebugPrinterMutex       sync.RWMutex
	setDebugPrinterArgsForCall []struct {
		arg1 loggregator_consumer.DebugPrinter
	}
}

func (fake *FakeLoggregatorConsumer) Tail(appGUID string, authToken string) (<-chan *logmessage.LogMessage, error) {
	fake.tailMutex.Lock()
	fake.tailArgsForCall = append(fake.tailArgsForCall, struct {
		appGUID   string
		authToken string
	}{appGUID, authToken})
	fake.tailMutex.Unlock()
	if fake.TailStub != nil {
		return fake.TailStub(appGUID, authToken)
	} else {
		return fake.tailReturns.result1, fake.tailReturns.result2
	}
}

func (fake *FakeLoggregatorConsumer) TailCallCount() int {
	fake.tailMutex.RLock()
	defer fake.tailMutex.RUnlock()
	return len(fake.tailArgsForCall)
}

func (fake *FakeLoggregatorConsumer) TailArgsForCall(i int) (string, string) {
	fake.tailMutex.RLock()
	defer fake.tailMutex.RUnlock()
	return fake.tailArgsForCall[i].appGUID, fake.tailArgsForCall[i].authToken
}

func (fake *FakeLoggregatorConsumer) TailReturns(result1 <-chan *logmessage.LogMessage, result2 error) {
	fake.TailStub = nil
	fake.tailReturns = struct {
		result1 <-chan *logmessage.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeLoggregatorConsumer) Recent(appGUID string, authToken string) ([]*logmessage.LogMessage, error) {
	fake.recentMutex.Lock()
	fake.recentArgsForCall = append(fake.recentArgsForCall, struct {
		appGUID   string
		authToken string
	}{appGUID, authToken})
	fake.recentMutex.Unlock()
	if fake.RecentStub != nil {
		return fake.RecentStub(appGUID, authToken)
	} else {
		return fake.recentReturns.result1, fake.recentReturns.result2
	}
}

func (fake *FakeLoggregatorConsumer) RecentCallCount() int {
	fake.recentMutex.RLock()
	defer fake.recentMutex.RUnlock()
	return len(fake.recentArgsForCall)
}

func (fake *FakeLoggregatorConsumer) RecentArgsForCall(i int) (string, string) {
	fake.recentMutex.RLock()
	defer fake.recentMutex.RUnlock()
	return fake.recentArgsForCall[i].appGUID, fake.recentArgsForCall[i].authToken
}

func (fake *FakeLoggregatorConsumer) RecentReturns(result1 []*logmessage.LogMessage, result2 error) {
	fake.RecentStub = nil
	fake.recentReturns = struct {
		result1 []*logmessage.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeLoggregatorConsumer) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeLoggregatorConsumer) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeLoggregatorConsumer) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLoggregatorConsumer) SetOnConnectCallback(arg1 func()) {
	fake.setOnConnectCallbackMutex.Lock()
	fake.setOnConnectCallbackArgsForCall = append(fake.setOnConnectCallbackArgsForCall, struct {
		arg1 func()
	}{arg1})
	fake.setOnConnectCallbackMutex.Unlock()
	if fake.SetOnConnectCallbackStub != nil {
		fake.SetOnConnectCallbackStub(arg1)
	}
}

func (fake *FakeLoggregatorConsumer) SetOnConnectCallbackCallCount() int {
	fake.setOnConnectCallbackMutex.RLock()
	defer fake.setOnConnectCallbackMutex.RUnlock()
	return len(fake.setOnConnectCallbackArgsForCall)
}

func (fake *FakeLoggregatorConsumer) SetOnConnectCallbackArgsForCall(i int) func() {
	fake.setOnConnectCallbackMutex.RLock()
	defer fake.setOnConnectCallbackMutex.RUnlock()
	return fake.setOnConnectCallbackArgsForCall[i].arg1
}

func (fake *FakeLoggregatorConsumer) SetDebugPrinter(arg1 loggregator_consumer.DebugPrinter) {
	fake.setDebugPrinterMutex.Lock()
	fake.setDebugPrinterArgsForCall = append(fake.setDebugPrinterArgsForCall, struct {
		arg1 loggregator_consumer.DebugPrinter
	}{arg1})
	fake.setDebugPrinterMutex.Unlock()
	if fake.SetDebugPrinterStub != nil {
		fake.SetDebugPrinterStub(arg1)
	}
}

func (fake *FakeLoggregatorConsumer) SetDebugPrinterCallCount() int {
	fake.setDebugPrinterMutex.RLock()
	defer fake.setDebugPrinterMutex.RUnlock()
	return len(fake.setDebugPrinterArgsForCall)
}

func (fake *FakeLoggregatorConsumer) SetDebugPrinterArgsForCall(i int) loggregator_consumer.DebugPrinter {
	fake.setDebugPrinterMutex.RLock()
	defer fake.setDebugPrinterMutex.RUnlock()
	return fake.setDebugPrinterArgsForCall[i].arg1
}

var _ logs.LoggregatorConsumer = new(FakeLoggregatorConsumer)
