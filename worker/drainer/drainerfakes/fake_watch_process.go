// Code generated by counterfeiter. DO NOT EDIT.
package drainerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/worker/drainer"
)

type FakeWatchProcess struct {
	IsRunningStub        func(lager.Logger) (bool, error)
	isRunningMutex       sync.RWMutex
	isRunningArgsForCall []struct {
		arg1 lager.Logger
	}
	isRunningReturns struct {
		result1 bool
		result2 error
	}
	isRunningReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWatchProcess) IsRunning(arg1 lager.Logger) (bool, error) {
	fake.isRunningMutex.Lock()
	ret, specificReturn := fake.isRunningReturnsOnCall[len(fake.isRunningArgsForCall)]
	fake.isRunningArgsForCall = append(fake.isRunningArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("IsRunning", []interface{}{arg1})
	fake.isRunningMutex.Unlock()
	if fake.IsRunningStub != nil {
		return fake.IsRunningStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isRunningReturns.result1, fake.isRunningReturns.result2
}

func (fake *FakeWatchProcess) IsRunningCallCount() int {
	fake.isRunningMutex.RLock()
	defer fake.isRunningMutex.RUnlock()
	return len(fake.isRunningArgsForCall)
}

func (fake *FakeWatchProcess) IsRunningArgsForCall(i int) lager.Logger {
	fake.isRunningMutex.RLock()
	defer fake.isRunningMutex.RUnlock()
	return fake.isRunningArgsForCall[i].arg1
}

func (fake *FakeWatchProcess) IsRunningReturns(result1 bool, result2 error) {
	fake.IsRunningStub = nil
	fake.isRunningReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeWatchProcess) IsRunningReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsRunningStub = nil
	if fake.isRunningReturnsOnCall == nil {
		fake.isRunningReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isRunningReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeWatchProcess) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isRunningMutex.RLock()
	defer fake.isRunningMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWatchProcess) recordInvocation(key string, args []interface{}) {
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

var _ drainer.WatchProcess = new(FakeWatchProcess)
