// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-park-mail-ru/2020_2_CodeExpress/internal/playlist (interfaces: PlaylistRep,PlaylistUsecase)

// Package mock_playlist is a generated GoMock package.
package mock_playlist

import (
	models "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/models"
	error_response "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/tools/error_response"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPlaylistRep is a mock of PlaylistRep interface
type MockPlaylistRep struct {
	ctrl     *gomock.Controller
	recorder *MockPlaylistRepMockRecorder
}

// MockPlaylistRepMockRecorder is the mock recorder for MockPlaylistRep
type MockPlaylistRepMockRecorder struct {
	mock *MockPlaylistRep
}

// NewMockPlaylistRep creates a new mock instance
func NewMockPlaylistRep(ctrl *gomock.Controller) *MockPlaylistRep {
	mock := &MockPlaylistRep{ctrl: ctrl}
	mock.recorder = &MockPlaylistRepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlaylistRep) EXPECT() *MockPlaylistRepMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockPlaylistRep) Delete(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPlaylistRepMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPlaylistRep)(nil).Delete), arg0)
}

// DeleteTrack mocks base method
func (m *MockPlaylistRep) DeleteTrack(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrack indicates an expected call of DeleteTrack
func (mr *MockPlaylistRepMockRecorder) DeleteTrack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrack", reflect.TypeOf((*MockPlaylistRep)(nil).DeleteTrack), arg0, arg1)
}

// Insert mocks base method
func (m *MockPlaylistRep) Insert(arg0 *models.Playlist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockPlaylistRepMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPlaylistRep)(nil).Insert), arg0)
}

// InsertTrack mocks base method
func (m *MockPlaylistRep) InsertTrack(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTrack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTrack indicates an expected call of InsertTrack
func (mr *MockPlaylistRepMockRecorder) InsertTrack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTrack", reflect.TypeOf((*MockPlaylistRep)(nil).InsertTrack), arg0, arg1)
}

// SelectByID mocks base method
func (m *MockPlaylistRep) SelectByID(arg0 uint64) (*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByID", arg0)
	ret0, _ := ret[0].(*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByID indicates an expected call of SelectByID
func (mr *MockPlaylistRepMockRecorder) SelectByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByID", reflect.TypeOf((*MockPlaylistRep)(nil).SelectByID), arg0)
}

// SelectByUserID mocks base method
func (m *MockPlaylistRep) SelectByUserID(arg0 uint64) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectByUserID", arg0)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByUserID indicates an expected call of SelectByUserID
func (mr *MockPlaylistRepMockRecorder) SelectByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByUserID", reflect.TypeOf((*MockPlaylistRep)(nil).SelectByUserID), arg0)
}

// SelectPublicByUserID mocks base method
func (m *MockPlaylistRep) SelectPublicByUserID(arg0 uint64) ([]*models.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectPublicByUserID", arg0)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectPublicByUserID indicates an expected call of SelectPublicByUserID
func (mr *MockPlaylistRepMockRecorder) SelectPublicByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectPublicByUserID", reflect.TypeOf((*MockPlaylistRep)(nil).SelectPublicByUserID), arg0)
}

// Update mocks base method
func (m *MockPlaylistRep) Update(arg0 *models.Playlist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockPlaylistRepMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPlaylistRep)(nil).Update), arg0)
}

// MockPlaylistUsecase is a mock of PlaylistUsecase interface
type MockPlaylistUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPlaylistUsecaseMockRecorder
}

// MockPlaylistUsecaseMockRecorder is the mock recorder for MockPlaylistUsecase
type MockPlaylistUsecaseMockRecorder struct {
	mock *MockPlaylistUsecase
}

// NewMockPlaylistUsecase creates a new mock instance
func NewMockPlaylistUsecase(ctrl *gomock.Controller) *MockPlaylistUsecase {
	mock := &MockPlaylistUsecase{ctrl: ctrl}
	mock.recorder = &MockPlaylistUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlaylistUsecase) EXPECT() *MockPlaylistUsecaseMockRecorder {
	return m.recorder
}

// AddTrack mocks base method
func (m *MockPlaylistUsecase) AddTrack(arg0, arg1 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTrack", arg0, arg1)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// AddTrack indicates an expected call of AddTrack
func (mr *MockPlaylistUsecaseMockRecorder) AddTrack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTrack", reflect.TypeOf((*MockPlaylistUsecase)(nil).AddTrack), arg0, arg1)
}

// CreatePlaylist mocks base method
func (m *MockPlaylistUsecase) CreatePlaylist(arg0 *models.Playlist) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaylist", arg0)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// CreatePlaylist indicates an expected call of CreatePlaylist
func (mr *MockPlaylistUsecaseMockRecorder) CreatePlaylist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaylist", reflect.TypeOf((*MockPlaylistUsecase)(nil).CreatePlaylist), arg0)
}

// DeletePlaylist mocks base method
func (m *MockPlaylistUsecase) DeletePlaylist(arg0 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlaylist", arg0)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// DeletePlaylist indicates an expected call of DeletePlaylist
func (mr *MockPlaylistUsecaseMockRecorder) DeletePlaylist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaylist", reflect.TypeOf((*MockPlaylistUsecase)(nil).DeletePlaylist), arg0)
}

// DeleteTrack mocks base method
func (m *MockPlaylistUsecase) DeleteTrack(arg0, arg1 uint64) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrack", arg0, arg1)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// DeleteTrack indicates an expected call of DeleteTrack
func (mr *MockPlaylistUsecaseMockRecorder) DeleteTrack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrack", reflect.TypeOf((*MockPlaylistUsecase)(nil).DeleteTrack), arg0, arg1)
}

// GetByID mocks base method
func (m *MockPlaylistUsecase) GetByID(arg0 uint64) (*models.Playlist, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*models.Playlist)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockPlaylistUsecaseMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPlaylistUsecase)(nil).GetByID), arg0)
}

// GetByUserID mocks base method
func (m *MockPlaylistUsecase) GetByUserID(arg0 uint64) ([]*models.Playlist, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserID", arg0)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetByUserID indicates an expected call of GetByUserID
func (mr *MockPlaylistUsecaseMockRecorder) GetByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockPlaylistUsecase)(nil).GetByUserID), arg0)
}

// GetPublicByUserID mocks base method
func (m *MockPlaylistUsecase) GetPublicByUserID(arg0 uint64) ([]*models.Playlist, *error_response.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicByUserID", arg0)
	ret0, _ := ret[0].([]*models.Playlist)
	ret1, _ := ret[1].(*error_response.ErrorResponse)
	return ret0, ret1
}

// GetPublicByUserID indicates an expected call of GetPublicByUserID
func (mr *MockPlaylistUsecaseMockRecorder) GetPublicByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicByUserID", reflect.TypeOf((*MockPlaylistUsecase)(nil).GetPublicByUserID), arg0)
}

// UpdatePlaylist mocks base method
func (m *MockPlaylistUsecase) UpdatePlaylist(arg0 *models.Playlist) *error_response.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlaylist", arg0)
	ret0, _ := ret[0].(*error_response.ErrorResponse)
	return ret0
}

// UpdatePlaylist indicates an expected call of UpdatePlaylist
func (mr *MockPlaylistUsecaseMockRecorder) UpdatePlaylist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlaylist", reflect.TypeOf((*MockPlaylistUsecase)(nil).UpdatePlaylist), arg0)
}
