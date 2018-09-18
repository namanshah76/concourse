// Code generated by counterfeiter. DO NOT EDIT.
package schedulerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/scheduler"
)

type FakeScanner struct {
	ScanStub        func(lager.Logger, string) error
	scanMutex       sync.RWMutex
	scanArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	scanReturns struct {
		result1 error
	}
	scanReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScanner) Scan(arg1 lager.Logger, arg2 string) error {
	fake.scanMutex.Lock()
	ret, specificReturn := fake.scanReturnsOnCall[len(fake.scanArgsForCall)]
	fake.scanArgsForCall = append(fake.scanArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Scan", []interface{}{arg1, arg2})
	fake.scanMutex.Unlock()
	if fake.ScanStub != nil {
		return fake.ScanStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.scanReturns.result1
}

func (fake *FakeScanner) ScanCallCount() int {
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	return len(fake.scanArgsForCall)
}

func (fake *FakeScanner) ScanArgsForCall(i int) (lager.Logger, string) {
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	return fake.scanArgsForCall[i].arg1, fake.scanArgsForCall[i].arg2
}

func (fake *FakeScanner) ScanReturns(result1 error) {
	fake.ScanStub = nil
	fake.scanReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScanner) ScanReturnsOnCall(i int, result1 error) {
	fake.ScanStub = nil
	if fake.scanReturnsOnCall == nil {
		fake.scanReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.scanReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScanner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeScanner) recordInvocation(key string, args []interface{}) {
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

var _ scheduler.Scanner = new(FakeScanner)
