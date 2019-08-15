// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package apiservice

import (
	"github.com/elim/GoCourses/Model"
	"sync"
)

var (
	lockRepoMockGetDetailData sync.RWMutex
)

// Ensure, that RepoMock does implement Repo.
// If this is not the case, regenerate this file with moq.
var _ Repo = &RepoMock{}

// RepoMock is a mock implementation of Repo.
//
//     func TestSomethingThatUsesRepo(t *testing.T) {
//
//         // make and configure a mocked Repo
//         mockedRepo := &RepoMock{
//             GetDetailDataFunc: func(id string) []model.DetailPosition {
// 	               panic("mock out the GetDetailData method")
//             },
//         }
//
//         // use mockedRepo in code that requires Repo
//         // and then make assertions.
//
//     }
type RepoMock struct {
	// GetDetailDataFunc mocks the GetDetailData method.
	GetDetailDataFunc func(id string) []model.DetailPosition

	// calls tracks calls to the methods.
	calls struct {
		// GetDetailData holds details about calls to the GetDetailData method.
		GetDetailData []struct {
			// ID is the id argument value.
			ID string
		}
	}
}

// GetDetailData calls GetDetailDataFunc.
func (mock *RepoMock) GetDetailData(id string) []model.DetailPosition {
	if mock.GetDetailDataFunc == nil {
		panic("RepoMock.GetDetailDataFunc: method is nil but Repo.GetDetailData was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	lockRepoMockGetDetailData.Lock()
	mock.calls.GetDetailData = append(mock.calls.GetDetailData, callInfo)
	lockRepoMockGetDetailData.Unlock()
	return mock.GetDetailDataFunc(id)
}

// GetDetailDataCalls gets all the calls that were made to GetDetailData.
// Check the length with:
//     len(mockedRepo.GetDetailDataCalls())
func (mock *RepoMock) GetDetailDataCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockRepoMockGetDetailData.RLock()
	calls = mock.calls.GetDetailData
	lockRepoMockGetDetailData.RUnlock()
	return calls
}