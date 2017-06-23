// Code generated by counterfeiter. DO NOT EDIT.
package v3actionfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/cli/actor/v3action"
)

type FakeConfig struct {
	PollingIntervalStub        func() time.Duration
	pollingIntervalMutex       sync.RWMutex
	pollingIntervalArgsForCall []struct{}
	pollingIntervalReturns     struct {
		result1 time.Duration
	}
	pollingIntervalReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	StartupTimeoutStub        func() time.Duration
	startupTimeoutMutex       sync.RWMutex
	startupTimeoutArgsForCall []struct{}
	startupTimeoutReturns     struct {
		result1 time.Duration
	}
	startupTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) PollingInterval() time.Duration {
	fake.pollingIntervalMutex.Lock()
	ret, specificReturn := fake.pollingIntervalReturnsOnCall[len(fake.pollingIntervalArgsForCall)]
	fake.pollingIntervalArgsForCall = append(fake.pollingIntervalArgsForCall, struct{}{})
	fake.recordInvocation("PollingInterval", []interface{}{})
	fake.pollingIntervalMutex.Unlock()
	if fake.PollingIntervalStub != nil {
		return fake.PollingIntervalStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pollingIntervalReturns.result1
}

func (fake *FakeConfig) PollingIntervalCallCount() int {
	fake.pollingIntervalMutex.RLock()
	defer fake.pollingIntervalMutex.RUnlock()
	return len(fake.pollingIntervalArgsForCall)
}

func (fake *FakeConfig) PollingIntervalReturns(result1 time.Duration) {
	fake.PollingIntervalStub = nil
	fake.pollingIntervalReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) PollingIntervalReturnsOnCall(i int, result1 time.Duration) {
	fake.PollingIntervalStub = nil
	if fake.pollingIntervalReturnsOnCall == nil {
		fake.pollingIntervalReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.pollingIntervalReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StartupTimeout() time.Duration {
	fake.startupTimeoutMutex.Lock()
	ret, specificReturn := fake.startupTimeoutReturnsOnCall[len(fake.startupTimeoutArgsForCall)]
	fake.startupTimeoutArgsForCall = append(fake.startupTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("StartupTimeout", []interface{}{})
	fake.startupTimeoutMutex.Unlock()
	if fake.StartupTimeoutStub != nil {
		return fake.StartupTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startupTimeoutReturns.result1
}

func (fake *FakeConfig) StartupTimeoutCallCount() int {
	fake.startupTimeoutMutex.RLock()
	defer fake.startupTimeoutMutex.RUnlock()
	return len(fake.startupTimeoutArgsForCall)
}

func (fake *FakeConfig) StartupTimeoutReturns(result1 time.Duration) {
	fake.StartupTimeoutStub = nil
	fake.startupTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StartupTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.StartupTimeoutStub = nil
	if fake.startupTimeoutReturnsOnCall == nil {
		fake.startupTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.startupTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pollingIntervalMutex.RLock()
	defer fake.pollingIntervalMutex.RUnlock()
	fake.startupTimeoutMutex.RLock()
	defer fake.startupTimeoutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
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

var _ v3action.Config = new(FakeConfig)
