// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeV3SetDropletActor struct {
	SetApplicationDropletStub        func(appName string, spaceGUID string, dropletGUID string) (v3action.Warnings, error)
	setApplicationDropletMutex       sync.RWMutex
	setApplicationDropletArgsForCall []struct {
		appName     string
		spaceGUID   string
		dropletGUID string
	}
	setApplicationDropletReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	setApplicationDropletReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3SetDropletActor) SetApplicationDroplet(appName string, spaceGUID string, dropletGUID string) (v3action.Warnings, error) {
	fake.setApplicationDropletMutex.Lock()
	ret, specificReturn := fake.setApplicationDropletReturnsOnCall[len(fake.setApplicationDropletArgsForCall)]
	fake.setApplicationDropletArgsForCall = append(fake.setApplicationDropletArgsForCall, struct {
		appName     string
		spaceGUID   string
		dropletGUID string
	}{appName, spaceGUID, dropletGUID})
	fake.recordInvocation("SetApplicationDroplet", []interface{}{appName, spaceGUID, dropletGUID})
	fake.setApplicationDropletMutex.Unlock()
	if fake.SetApplicationDropletStub != nil {
		return fake.SetApplicationDropletStub(appName, spaceGUID, dropletGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setApplicationDropletReturns.result1, fake.setApplicationDropletReturns.result2
}

func (fake *FakeV3SetDropletActor) SetApplicationDropletCallCount() int {
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	return len(fake.setApplicationDropletArgsForCall)
}

func (fake *FakeV3SetDropletActor) SetApplicationDropletArgsForCall(i int) (string, string, string) {
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	return fake.setApplicationDropletArgsForCall[i].appName, fake.setApplicationDropletArgsForCall[i].spaceGUID, fake.setApplicationDropletArgsForCall[i].dropletGUID
}

func (fake *FakeV3SetDropletActor) SetApplicationDropletReturns(result1 v3action.Warnings, result2 error) {
	fake.SetApplicationDropletStub = nil
	fake.setApplicationDropletReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3SetDropletActor) SetApplicationDropletReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.SetApplicationDropletStub = nil
	if fake.setApplicationDropletReturnsOnCall == nil {
		fake.setApplicationDropletReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.setApplicationDropletReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3SetDropletActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3SetDropletActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.V3SetDropletActor = new(FakeV3SetDropletActor)
