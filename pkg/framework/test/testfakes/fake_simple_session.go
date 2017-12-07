// Code generated by counterfeiter. DO NOT EDIT.
package testfakes

import (
	"sync"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"k8s.io/kubectl/pkg/framework/test"
)

type FakeSimpleSession struct {
	BufferStub        func() *gbytes.Buffer
	bufferMutex       sync.RWMutex
	bufferArgsForCall []struct{}
	bufferReturns     struct {
		result1 *gbytes.Buffer
	}
	bufferReturnsOnCall map[int]struct {
		result1 *gbytes.Buffer
	}
	ExitCodeStub        func() int
	exitCodeMutex       sync.RWMutex
	exitCodeArgsForCall []struct{}
	exitCodeReturns     struct {
		result1 int
	}
	exitCodeReturnsOnCall map[int]struct {
		result1 int
	}
	WaitStub        func(timeout ...interface{}) *gexec.Session
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
		timeout []interface{}
	}
	waitReturns struct {
		result1 *gexec.Session
	}
	waitReturnsOnCall map[int]struct {
		result1 *gexec.Session
	}
	TerminateStub        func() *gexec.Session
	terminateMutex       sync.RWMutex
	terminateArgsForCall []struct{}
	terminateReturns     struct {
		result1 *gexec.Session
	}
	terminateReturnsOnCall map[int]struct {
		result1 *gexec.Session
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSimpleSession) Buffer() *gbytes.Buffer {
	fake.bufferMutex.Lock()
	ret, specificReturn := fake.bufferReturnsOnCall[len(fake.bufferArgsForCall)]
	fake.bufferArgsForCall = append(fake.bufferArgsForCall, struct{}{})
	fake.recordInvocation("Buffer", []interface{}{})
	fake.bufferMutex.Unlock()
	if fake.BufferStub != nil {
		return fake.BufferStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.bufferReturns.result1
}

func (fake *FakeSimpleSession) BufferCallCount() int {
	fake.bufferMutex.RLock()
	defer fake.bufferMutex.RUnlock()
	return len(fake.bufferArgsForCall)
}

func (fake *FakeSimpleSession) BufferReturns(result1 *gbytes.Buffer) {
	fake.BufferStub = nil
	fake.bufferReturns = struct {
		result1 *gbytes.Buffer
	}{result1}
}

func (fake *FakeSimpleSession) BufferReturnsOnCall(i int, result1 *gbytes.Buffer) {
	fake.BufferStub = nil
	if fake.bufferReturnsOnCall == nil {
		fake.bufferReturnsOnCall = make(map[int]struct {
			result1 *gbytes.Buffer
		})
	}
	fake.bufferReturnsOnCall[i] = struct {
		result1 *gbytes.Buffer
	}{result1}
}

func (fake *FakeSimpleSession) ExitCode() int {
	fake.exitCodeMutex.Lock()
	ret, specificReturn := fake.exitCodeReturnsOnCall[len(fake.exitCodeArgsForCall)]
	fake.exitCodeArgsForCall = append(fake.exitCodeArgsForCall, struct{}{})
	fake.recordInvocation("ExitCode", []interface{}{})
	fake.exitCodeMutex.Unlock()
	if fake.ExitCodeStub != nil {
		return fake.ExitCodeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.exitCodeReturns.result1
}

func (fake *FakeSimpleSession) ExitCodeCallCount() int {
	fake.exitCodeMutex.RLock()
	defer fake.exitCodeMutex.RUnlock()
	return len(fake.exitCodeArgsForCall)
}

func (fake *FakeSimpleSession) ExitCodeReturns(result1 int) {
	fake.ExitCodeStub = nil
	fake.exitCodeReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeSimpleSession) ExitCodeReturnsOnCall(i int, result1 int) {
	fake.ExitCodeStub = nil
	if fake.exitCodeReturnsOnCall == nil {
		fake.exitCodeReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.exitCodeReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeSimpleSession) Wait(timeout ...interface{}) *gexec.Session {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
		timeout []interface{}
	}{timeout})
	fake.recordInvocation("Wait", []interface{}{timeout})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub(timeout...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.waitReturns.result1
}

func (fake *FakeSimpleSession) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeSimpleSession) WaitArgsForCall(i int) []interface{} {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return fake.waitArgsForCall[i].timeout
}

func (fake *FakeSimpleSession) WaitReturns(result1 *gexec.Session) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 *gexec.Session
	}{result1}
}

func (fake *FakeSimpleSession) WaitReturnsOnCall(i int, result1 *gexec.Session) {
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 *gexec.Session
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 *gexec.Session
	}{result1}
}

func (fake *FakeSimpleSession) Terminate() *gexec.Session {
	fake.terminateMutex.Lock()
	ret, specificReturn := fake.terminateReturnsOnCall[len(fake.terminateArgsForCall)]
	fake.terminateArgsForCall = append(fake.terminateArgsForCall, struct{}{})
	fake.recordInvocation("Terminate", []interface{}{})
	fake.terminateMutex.Unlock()
	if fake.TerminateStub != nil {
		return fake.TerminateStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.terminateReturns.result1
}

func (fake *FakeSimpleSession) TerminateCallCount() int {
	fake.terminateMutex.RLock()
	defer fake.terminateMutex.RUnlock()
	return len(fake.terminateArgsForCall)
}

func (fake *FakeSimpleSession) TerminateReturns(result1 *gexec.Session) {
	fake.TerminateStub = nil
	fake.terminateReturns = struct {
		result1 *gexec.Session
	}{result1}
}

func (fake *FakeSimpleSession) TerminateReturnsOnCall(i int, result1 *gexec.Session) {
	fake.TerminateStub = nil
	if fake.terminateReturnsOnCall == nil {
		fake.terminateReturnsOnCall = make(map[int]struct {
			result1 *gexec.Session
		})
	}
	fake.terminateReturnsOnCall[i] = struct {
		result1 *gexec.Session
	}{result1}
}

func (fake *FakeSimpleSession) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bufferMutex.RLock()
	defer fake.bufferMutex.RUnlock()
	fake.exitCodeMutex.RLock()
	defer fake.exitCodeMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	fake.terminateMutex.RLock()
	defer fake.terminateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSimpleSession) recordInvocation(key string, args []interface{}) {
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

var _ test.SimpleSession = new(FakeSimpleSession)