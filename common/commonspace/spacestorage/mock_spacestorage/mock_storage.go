// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/storage (interfaces: SpaceStorageProvider,SpaceStorage)

// Package mock_storage is a generated GoMock package.
package mock_spacestorage

import (
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/object/acl/liststorage"
	storage3 "github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/object/tree/treestorage"
	reflect "reflect"

	app "github.com/anytypeio/go-anytype-infrastructure-experiments/common/app"
	storage "github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/spacestorage"
	spacesyncproto "github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/spacesyncproto"
	gomock "github.com/golang/mock/gomock"
)

// MockSpaceStorageProvider is a mock of SpaceStorageProvider interface.
type MockSpaceStorageProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceStorageProviderMockRecorder
}

// MockSpaceStorageProviderMockRecorder is the mock recorder for MockSpaceStorageProvider.
type MockSpaceStorageProviderMockRecorder struct {
	mock *MockSpaceStorageProvider
}

// NewMockSpaceStorageProvider creates a new mock instance.
func NewMockSpaceStorageProvider(ctrl *gomock.Controller) *MockSpaceStorageProvider {
	mock := &MockSpaceStorageProvider{ctrl: ctrl}
	mock.recorder = &MockSpaceStorageProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceStorageProvider) EXPECT() *MockSpaceStorageProviderMockRecorder {
	return m.recorder
}

// CreateSpaceStorage mocks base method.
func (m *MockSpaceStorageProvider) CreateSpaceStorage(arg0 storage.SpaceStorageCreatePayload) (storage.SpaceStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSpaceStorage", arg0)
	ret0, _ := ret[0].(storage.SpaceStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSpaceStorage indicates an expected call of CreateSpaceStorage.
func (mr *MockSpaceStorageProviderMockRecorder) CreateSpaceStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSpaceStorage", reflect.TypeOf((*MockSpaceStorageProvider)(nil).CreateSpaceStorage), arg0)
}

// Init mocks base method.
func (m *MockSpaceStorageProvider) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockSpaceStorageProviderMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSpaceStorageProvider)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockSpaceStorageProvider) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockSpaceStorageProviderMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockSpaceStorageProvider)(nil).Name))
}

// SpaceStorage mocks base method.
func (m *MockSpaceStorageProvider) SpaceStorage(arg0 string) (storage.SpaceStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceStorage", arg0)
	ret0, _ := ret[0].(storage.SpaceStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceStorage indicates an expected call of SpaceStorage.
func (mr *MockSpaceStorageProviderMockRecorder) SpaceStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceStorage", reflect.TypeOf((*MockSpaceStorageProvider)(nil).SpaceStorage), arg0)
}

// MockSpaceStorage is a mock of SpaceStorage interface.
type MockSpaceStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceStorageMockRecorder
}

// MockSpaceStorageMockRecorder is the mock recorder for MockSpaceStorage.
type MockSpaceStorageMockRecorder struct {
	mock *MockSpaceStorage
}

// NewMockSpaceStorage creates a new mock instance.
func NewMockSpaceStorage(ctrl *gomock.Controller) *MockSpaceStorage {
	mock := &MockSpaceStorage{ctrl: ctrl}
	mock.recorder = &MockSpaceStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceStorage) EXPECT() *MockSpaceStorageMockRecorder {
	return m.recorder
}

// AclStorage mocks base method.
func (m *MockSpaceStorage) AclStorage() (liststorage.ListStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AclStorage")
	ret0, _ := ret[0].(liststorage.ListStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AclStorage indicates an expected call of AclStorage.
func (mr *MockSpaceStorageMockRecorder) AclStorage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AclStorage", reflect.TypeOf((*MockSpaceStorage)(nil).AclStorage))
}

// Close mocks base method.
func (m *MockSpaceStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSpaceStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSpaceStorage)(nil).Close))
}

// CreateTreeStorage mocks base method.
func (m *MockSpaceStorage) CreateTreeStorage(arg0 storage3.TreeStorageCreatePayload) (storage3.TreeStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTreeStorage", arg0)
	ret0, _ := ret[0].(storage3.TreeStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTreeStorage indicates an expected call of CreateTreeStorage.
func (mr *MockSpaceStorageMockRecorder) CreateTreeStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTreeStorage", reflect.TypeOf((*MockSpaceStorage)(nil).CreateTreeStorage), arg0)
}

// Id mocks base method.
func (m *MockSpaceStorage) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockSpaceStorageMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockSpaceStorage)(nil).Id))
}

// SetTreeDeletedStatus mocks base method.
func (m *MockSpaceStorage) SetTreeDeletedStatus(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTreeDeletedStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTreeDeletedStatus indicates an expected call of SetTreeDeletedStatus.
func (mr *MockSpaceStorageMockRecorder) SetTreeDeletedStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTreeDeletedStatus", reflect.TypeOf((*MockSpaceStorage)(nil).SetTreeDeletedStatus), arg0, arg1)
}

// SpaceHeader mocks base method.
func (m *MockSpaceStorage) SpaceHeader() (*spacesyncproto.RawSpaceHeaderWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceHeader")
	ret0, _ := ret[0].(*spacesyncproto.RawSpaceHeaderWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceHeader indicates an expected call of SpaceHeader.
func (mr *MockSpaceStorageMockRecorder) SpaceHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceHeader", reflect.TypeOf((*MockSpaceStorage)(nil).SpaceHeader))
}

// SpaceSettingsId mocks base method.
func (m *MockSpaceStorage) SpaceSettingsId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceSettingsId")
	ret0, _ := ret[0].(string)
	return ret0
}

// SpaceSettingsId indicates an expected call of SpaceSettingsId.
func (mr *MockSpaceStorageMockRecorder) SpaceSettingsId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceSettingsId", reflect.TypeOf((*MockSpaceStorage)(nil).SpaceSettingsId))
}

// StoredIds mocks base method.
func (m *MockSpaceStorage) StoredIds() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoredIds")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoredIds indicates an expected call of StoredIds.
func (mr *MockSpaceStorageMockRecorder) StoredIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoredIds", reflect.TypeOf((*MockSpaceStorage)(nil).StoredIds))
}

// TreeDeletedStatus mocks base method.
func (m *MockSpaceStorage) TreeDeletedStatus(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TreeDeletedStatus", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TreeDeletedStatus indicates an expected call of TreeDeletedStatus.
func (mr *MockSpaceStorageMockRecorder) TreeDeletedStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TreeDeletedStatus", reflect.TypeOf((*MockSpaceStorage)(nil).TreeDeletedStatus), arg0)
}

// TreeStorage mocks base method.
func (m *MockSpaceStorage) TreeStorage(arg0 string) (storage3.TreeStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TreeStorage", arg0)
	ret0, _ := ret[0].(storage3.TreeStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TreeStorage indicates an expected call of TreeStorage.
func (mr *MockSpaceStorageMockRecorder) TreeStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TreeStorage", reflect.TypeOf((*MockSpaceStorage)(nil).TreeStorage), arg0)
}
