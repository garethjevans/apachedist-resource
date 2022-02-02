// Code generated by counterfeiter. DO NOT EDIT.
package downloadfakes

import (
	"sync"

	semver "github.com/Masterminds/semver/v3"
	"github.com/garethjevans/apachedist-resource/download"
)

type FakeDownloader struct {
	DownloadStub        func(string, string, string, string, string) (*download.DownloadedArtifact, error)
	downloadMutex       sync.RWMutex
	downloadArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	downloadReturns struct {
		result1 *download.DownloadedArtifact
		result2 error
	}
	downloadReturnsOnCall map[int]struct {
		result1 *download.DownloadedArtifact
		result2 error
	}
	GetVersionsStub        func(string) ([]*semver.Version, error)
	getVersionsMutex       sync.RWMutex
	getVersionsArgsForCall []struct {
		arg1 string
	}
	getVersionsReturns struct {
		result1 []*semver.Version
		result2 error
	}
	getVersionsReturnsOnCall map[int]struct {
		result1 []*semver.Version
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDownloader) Download(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) (*download.DownloadedArtifact, error) {
	fake.downloadMutex.Lock()
	ret, specificReturn := fake.downloadReturnsOnCall[len(fake.downloadArgsForCall)]
	fake.downloadArgsForCall = append(fake.downloadArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.DownloadStub
	fakeReturns := fake.downloadReturns
	fake.recordInvocation("Download", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.downloadMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDownloader) DownloadCallCount() int {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	return len(fake.downloadArgsForCall)
}

func (fake *FakeDownloader) DownloadCalls(stub func(string, string, string, string, string) (*download.DownloadedArtifact, error)) {
	fake.downloadMutex.Lock()
	defer fake.downloadMutex.Unlock()
	fake.DownloadStub = stub
}

func (fake *FakeDownloader) DownloadArgsForCall(i int) (string, string, string, string, string) {
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	argsForCall := fake.downloadArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeDownloader) DownloadReturns(result1 *download.DownloadedArtifact, result2 error) {
	fake.downloadMutex.Lock()
	defer fake.downloadMutex.Unlock()
	fake.DownloadStub = nil
	fake.downloadReturns = struct {
		result1 *download.DownloadedArtifact
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloader) DownloadReturnsOnCall(i int, result1 *download.DownloadedArtifact, result2 error) {
	fake.downloadMutex.Lock()
	defer fake.downloadMutex.Unlock()
	fake.DownloadStub = nil
	if fake.downloadReturnsOnCall == nil {
		fake.downloadReturnsOnCall = make(map[int]struct {
			result1 *download.DownloadedArtifact
			result2 error
		})
	}
	fake.downloadReturnsOnCall[i] = struct {
		result1 *download.DownloadedArtifact
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloader) GetVersions(arg1 string) ([]*semver.Version, error) {
	fake.getVersionsMutex.Lock()
	ret, specificReturn := fake.getVersionsReturnsOnCall[len(fake.getVersionsArgsForCall)]
	fake.getVersionsArgsForCall = append(fake.getVersionsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetVersionsStub
	fakeReturns := fake.getVersionsReturns
	fake.recordInvocation("GetVersions", []interface{}{arg1})
	fake.getVersionsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDownloader) GetVersionsCallCount() int {
	fake.getVersionsMutex.RLock()
	defer fake.getVersionsMutex.RUnlock()
	return len(fake.getVersionsArgsForCall)
}

func (fake *FakeDownloader) GetVersionsCalls(stub func(string) ([]*semver.Version, error)) {
	fake.getVersionsMutex.Lock()
	defer fake.getVersionsMutex.Unlock()
	fake.GetVersionsStub = stub
}

func (fake *FakeDownloader) GetVersionsArgsForCall(i int) string {
	fake.getVersionsMutex.RLock()
	defer fake.getVersionsMutex.RUnlock()
	argsForCall := fake.getVersionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDownloader) GetVersionsReturns(result1 []*semver.Version, result2 error) {
	fake.getVersionsMutex.Lock()
	defer fake.getVersionsMutex.Unlock()
	fake.GetVersionsStub = nil
	fake.getVersionsReturns = struct {
		result1 []*semver.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloader) GetVersionsReturnsOnCall(i int, result1 []*semver.Version, result2 error) {
	fake.getVersionsMutex.Lock()
	defer fake.getVersionsMutex.Unlock()
	fake.GetVersionsStub = nil
	if fake.getVersionsReturnsOnCall == nil {
		fake.getVersionsReturnsOnCall = make(map[int]struct {
			result1 []*semver.Version
			result2 error
		})
	}
	fake.getVersionsReturnsOnCall[i] = struct {
		result1 []*semver.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeDownloader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadMutex.RLock()
	defer fake.downloadMutex.RUnlock()
	fake.getVersionsMutex.RLock()
	defer fake.getVersionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDownloader) recordInvocation(key string, args []interface{}) {
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

var _ download.Downloader = new(FakeDownloader)
