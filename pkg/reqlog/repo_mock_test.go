// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package reqlog_test

import (
	"context"
	"github.com/dstotijn/hetty/pkg/reqlog"
	"github.com/dstotijn/hetty/pkg/scope"
	"net/http"
	"sync"
	"time"
)

// Ensure, that RepoMock does implement reqlog.Repository.
// If this is not the case, regenerate this file with moq.
var _ reqlog.Repository = &RepoMock{}

// RepoMock is a mock implementation of reqlog.Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked reqlog.Repository
// 		mockedRepository := &RepoMock{
// 			AddRequestLogFunc: func(ctx context.Context, req http.Request, body []byte, timestamp time.Time) (*reqlog.Request, error) {
// 				panic("mock out the AddRequestLog method")
// 			},
// 			AddResponseLogFunc: func(ctx context.Context, reqID int64, res http.Response, body []byte, timestamp time.Time) (*reqlog.Response, error) {
// 				panic("mock out the AddResponseLog method")
// 			},
// 			ClearRequestLogsFunc: func(ctx context.Context) error {
// 				panic("mock out the ClearRequestLogs method")
// 			},
// 			FindRequestLogByIDFunc: func(ctx context.Context, id int64) (reqlog.Request, error) {
// 				panic("mock out the FindRequestLogByID method")
// 			},
// 			FindRequestLogsFunc: func(ctx context.Context, filter reqlog.FindRequestsFilter, scopeMoqParam *scope.Scope) ([]reqlog.Request, error) {
// 				panic("mock out the FindRequestLogs method")
// 			},
// 			FindSettingsByModuleFunc: func(ctx context.Context, module string, settings interface{}) error {
// 				panic("mock out the FindSettingsByModule method")
// 			},
// 			UpsertSettingsFunc: func(ctx context.Context, module string, settings interface{}) error {
// 				panic("mock out the UpsertSettings method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires reqlog.Repository
// 		// and then make assertions.
//
// 	}
type RepoMock struct {
	// AddRequestLogFunc mocks the AddRequestLog method.
	AddRequestLogFunc func(ctx context.Context, req http.Request, body []byte, timestamp time.Time) (*reqlog.Request, error)

	// AddResponseLogFunc mocks the AddResponseLog method.
	AddResponseLogFunc func(ctx context.Context, reqID int64, res http.Response, body []byte, timestamp time.Time) (*reqlog.Response, error)

	// ClearRequestLogsFunc mocks the ClearRequestLogs method.
	ClearRequestLogsFunc func(ctx context.Context) error

	// FindRequestLogByIDFunc mocks the FindRequestLogByID method.
	FindRequestLogByIDFunc func(ctx context.Context, id int64) (reqlog.Request, error)

	// FindRequestLogsFunc mocks the FindRequestLogs method.
	FindRequestLogsFunc func(ctx context.Context, filter reqlog.FindRequestsFilter, scopeMoqParam *scope.Scope) ([]reqlog.Request, error)

	// FindSettingsByModuleFunc mocks the FindSettingsByModule method.
	FindSettingsByModuleFunc func(ctx context.Context, module string, settings interface{}) error

	// UpsertSettingsFunc mocks the UpsertSettings method.
	UpsertSettingsFunc func(ctx context.Context, module string, settings interface{}) error

	// calls tracks calls to the methods.
	calls struct {
		// AddRequestLog holds details about calls to the AddRequestLog method.
		AddRequestLog []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req http.Request
			// Body is the body argument value.
			Body []byte
			// Timestamp is the timestamp argument value.
			Timestamp time.Time
		}
		// AddResponseLog holds details about calls to the AddResponseLog method.
		AddResponseLog []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReqID is the reqID argument value.
			ReqID int64
			// Res is the res argument value.
			Res http.Response
			// Body is the body argument value.
			Body []byte
			// Timestamp is the timestamp argument value.
			Timestamp time.Time
		}
		// ClearRequestLogs holds details about calls to the ClearRequestLogs method.
		ClearRequestLogs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// FindRequestLogByID holds details about calls to the FindRequestLogByID method.
		FindRequestLogByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int64
		}
		// FindRequestLogs holds details about calls to the FindRequestLogs method.
		FindRequestLogs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter reqlog.FindRequestsFilter
			// ScopeMoqParam is the scopeMoqParam argument value.
			ScopeMoqParam *scope.Scope
		}
		// FindSettingsByModule holds details about calls to the FindSettingsByModule method.
		FindSettingsByModule []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Module is the module argument value.
			Module string
			// Settings is the settings argument value.
			Settings interface{}
		}
		// UpsertSettings holds details about calls to the UpsertSettings method.
		UpsertSettings []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Module is the module argument value.
			Module string
			// Settings is the settings argument value.
			Settings interface{}
		}
	}
	lockAddRequestLog        sync.RWMutex
	lockAddResponseLog       sync.RWMutex
	lockClearRequestLogs     sync.RWMutex
	lockFindRequestLogByID   sync.RWMutex
	lockFindRequestLogs      sync.RWMutex
	lockFindSettingsByModule sync.RWMutex
	lockUpsertSettings       sync.RWMutex
}

// AddRequestLog calls AddRequestLogFunc.
func (mock *RepoMock) AddRequestLog(ctx context.Context, req http.Request, body []byte, timestamp time.Time) (*reqlog.Request, error) {
	if mock.AddRequestLogFunc == nil {
		panic("RepoMock.AddRequestLogFunc: method is nil but Repository.AddRequestLog was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Req       http.Request
		Body      []byte
		Timestamp time.Time
	}{
		Ctx:       ctx,
		Req:       req,
		Body:      body,
		Timestamp: timestamp,
	}
	mock.lockAddRequestLog.Lock()
	mock.calls.AddRequestLog = append(mock.calls.AddRequestLog, callInfo)
	mock.lockAddRequestLog.Unlock()
	return mock.AddRequestLogFunc(ctx, req, body, timestamp)
}

// AddRequestLogCalls gets all the calls that were made to AddRequestLog.
// Check the length with:
//     len(mockedRepository.AddRequestLogCalls())
func (mock *RepoMock) AddRequestLogCalls() []struct {
	Ctx       context.Context
	Req       http.Request
	Body      []byte
	Timestamp time.Time
} {
	var calls []struct {
		Ctx       context.Context
		Req       http.Request
		Body      []byte
		Timestamp time.Time
	}
	mock.lockAddRequestLog.RLock()
	calls = mock.calls.AddRequestLog
	mock.lockAddRequestLog.RUnlock()
	return calls
}

// AddResponseLog calls AddResponseLogFunc.
func (mock *RepoMock) AddResponseLog(ctx context.Context, reqID int64, res http.Response, body []byte, timestamp time.Time) (*reqlog.Response, error) {
	if mock.AddResponseLogFunc == nil {
		panic("RepoMock.AddResponseLogFunc: method is nil but Repository.AddResponseLog was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		ReqID     int64
		Res       http.Response
		Body      []byte
		Timestamp time.Time
	}{
		Ctx:       ctx,
		ReqID:     reqID,
		Res:       res,
		Body:      body,
		Timestamp: timestamp,
	}
	mock.lockAddResponseLog.Lock()
	mock.calls.AddResponseLog = append(mock.calls.AddResponseLog, callInfo)
	mock.lockAddResponseLog.Unlock()
	return mock.AddResponseLogFunc(ctx, reqID, res, body, timestamp)
}

// AddResponseLogCalls gets all the calls that were made to AddResponseLog.
// Check the length with:
//     len(mockedRepository.AddResponseLogCalls())
func (mock *RepoMock) AddResponseLogCalls() []struct {
	Ctx       context.Context
	ReqID     int64
	Res       http.Response
	Body      []byte
	Timestamp time.Time
} {
	var calls []struct {
		Ctx       context.Context
		ReqID     int64
		Res       http.Response
		Body      []byte
		Timestamp time.Time
	}
	mock.lockAddResponseLog.RLock()
	calls = mock.calls.AddResponseLog
	mock.lockAddResponseLog.RUnlock()
	return calls
}

// ClearRequestLogs calls ClearRequestLogsFunc.
func (mock *RepoMock) ClearRequestLogs(ctx context.Context) error {
	if mock.ClearRequestLogsFunc == nil {
		panic("RepoMock.ClearRequestLogsFunc: method is nil but Repository.ClearRequestLogs was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockClearRequestLogs.Lock()
	mock.calls.ClearRequestLogs = append(mock.calls.ClearRequestLogs, callInfo)
	mock.lockClearRequestLogs.Unlock()
	return mock.ClearRequestLogsFunc(ctx)
}

// ClearRequestLogsCalls gets all the calls that were made to ClearRequestLogs.
// Check the length with:
//     len(mockedRepository.ClearRequestLogsCalls())
func (mock *RepoMock) ClearRequestLogsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockClearRequestLogs.RLock()
	calls = mock.calls.ClearRequestLogs
	mock.lockClearRequestLogs.RUnlock()
	return calls
}

// FindRequestLogByID calls FindRequestLogByIDFunc.
func (mock *RepoMock) FindRequestLogByID(ctx context.Context, id int64) (reqlog.Request, error) {
	if mock.FindRequestLogByIDFunc == nil {
		panic("RepoMock.FindRequestLogByIDFunc: method is nil but Repository.FindRequestLogByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int64
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFindRequestLogByID.Lock()
	mock.calls.FindRequestLogByID = append(mock.calls.FindRequestLogByID, callInfo)
	mock.lockFindRequestLogByID.Unlock()
	return mock.FindRequestLogByIDFunc(ctx, id)
}

// FindRequestLogByIDCalls gets all the calls that were made to FindRequestLogByID.
// Check the length with:
//     len(mockedRepository.FindRequestLogByIDCalls())
func (mock *RepoMock) FindRequestLogByIDCalls() []struct {
	Ctx context.Context
	ID  int64
} {
	var calls []struct {
		Ctx context.Context
		ID  int64
	}
	mock.lockFindRequestLogByID.RLock()
	calls = mock.calls.FindRequestLogByID
	mock.lockFindRequestLogByID.RUnlock()
	return calls
}

// FindRequestLogs calls FindRequestLogsFunc.
func (mock *RepoMock) FindRequestLogs(ctx context.Context, filter reqlog.FindRequestsFilter, scopeMoqParam *scope.Scope) ([]reqlog.Request, error) {
	if mock.FindRequestLogsFunc == nil {
		panic("RepoMock.FindRequestLogsFunc: method is nil but Repository.FindRequestLogs was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		Filter        reqlog.FindRequestsFilter
		ScopeMoqParam *scope.Scope
	}{
		Ctx:           ctx,
		Filter:        filter,
		ScopeMoqParam: scopeMoqParam,
	}
	mock.lockFindRequestLogs.Lock()
	mock.calls.FindRequestLogs = append(mock.calls.FindRequestLogs, callInfo)
	mock.lockFindRequestLogs.Unlock()
	return mock.FindRequestLogsFunc(ctx, filter, scopeMoqParam)
}

// FindRequestLogsCalls gets all the calls that were made to FindRequestLogs.
// Check the length with:
//     len(mockedRepository.FindRequestLogsCalls())
func (mock *RepoMock) FindRequestLogsCalls() []struct {
	Ctx           context.Context
	Filter        reqlog.FindRequestsFilter
	ScopeMoqParam *scope.Scope
} {
	var calls []struct {
		Ctx           context.Context
		Filter        reqlog.FindRequestsFilter
		ScopeMoqParam *scope.Scope
	}
	mock.lockFindRequestLogs.RLock()
	calls = mock.calls.FindRequestLogs
	mock.lockFindRequestLogs.RUnlock()
	return calls
}

// FindSettingsByModule calls FindSettingsByModuleFunc.
func (mock *RepoMock) FindSettingsByModule(ctx context.Context, module string, settings interface{}) error {
	if mock.FindSettingsByModuleFunc == nil {
		panic("RepoMock.FindSettingsByModuleFunc: method is nil but Repository.FindSettingsByModule was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Module   string
		Settings interface{}
	}{
		Ctx:      ctx,
		Module:   module,
		Settings: settings,
	}
	mock.lockFindSettingsByModule.Lock()
	mock.calls.FindSettingsByModule = append(mock.calls.FindSettingsByModule, callInfo)
	mock.lockFindSettingsByModule.Unlock()
	return mock.FindSettingsByModuleFunc(ctx, module, settings)
}

// FindSettingsByModuleCalls gets all the calls that were made to FindSettingsByModule.
// Check the length with:
//     len(mockedRepository.FindSettingsByModuleCalls())
func (mock *RepoMock) FindSettingsByModuleCalls() []struct {
	Ctx      context.Context
	Module   string
	Settings interface{}
} {
	var calls []struct {
		Ctx      context.Context
		Module   string
		Settings interface{}
	}
	mock.lockFindSettingsByModule.RLock()
	calls = mock.calls.FindSettingsByModule
	mock.lockFindSettingsByModule.RUnlock()
	return calls
}

// UpsertSettings calls UpsertSettingsFunc.
func (mock *RepoMock) UpsertSettings(ctx context.Context, module string, settings interface{}) error {
	if mock.UpsertSettingsFunc == nil {
		panic("RepoMock.UpsertSettingsFunc: method is nil but Repository.UpsertSettings was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Module   string
		Settings interface{}
	}{
		Ctx:      ctx,
		Module:   module,
		Settings: settings,
	}
	mock.lockUpsertSettings.Lock()
	mock.calls.UpsertSettings = append(mock.calls.UpsertSettings, callInfo)
	mock.lockUpsertSettings.Unlock()
	return mock.UpsertSettingsFunc(ctx, module, settings)
}

// UpsertSettingsCalls gets all the calls that were made to UpsertSettings.
// Check the length with:
//     len(mockedRepository.UpsertSettingsCalls())
func (mock *RepoMock) UpsertSettingsCalls() []struct {
	Ctx      context.Context
	Module   string
	Settings interface{}
} {
	var calls []struct {
		Ctx      context.Context
		Module   string
		Settings interface{}
	}
	mock.lockUpsertSettings.RLock()
	calls = mock.calls.UpsertSettings
	mock.lockUpsertSettings.RUnlock()
	return calls
}