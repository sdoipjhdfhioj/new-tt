package handler

import (
	"awesomeProject1/internal/handler/dto"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type nameServiceMock struct {
	mock.Mock
}

func (m *nameServiceMock) Set(nameRequest dto.SetNameRequest) error {

	args := m.Called(nameRequest)
	return args.Error(0)

}

func (m *nameServiceMock) Get(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *nameServiceMock) Remove(id string) error {

	args := m.Called(id)
	return args.Error(0)

}

func TestNameHandler(t *testing.T) {

	nameService := new(nameServiceMock)
	nh := NameHandler{
		NameManager: nameService,
	}

	NameHandler_GetName_Success(t, nameService, nh)
	NameHandler_GetName_NameGetError(t, nameService, nh)
	NameHandler_GetName_WithoutId(t, nameService, nh)

	NameHandler_SetName_Success(t, nameService, nh)
	NameHandler_SetName_WithoutIdOrName(t, nameService, nh)
	NameHandler_SetName_NameSetError(t, nameService, nh)

	NameHandler_RemoveName_Success(t, nameService, nh)
	NameHandler_RemoveName_WithoutId(t, nameService, nh)
	NameHandler_RemoveName_NameRemoveError(t, nameService, nh)

}

func NameHandler_GetName_Success(t *testing.T, nameService *nameServiceMock, nh NameHandler) {
	nameService.On("Get", "123").Return("name-123", nil)
	request, err := http.NewRequest(http.MethodGet, "/get/name?id=123", nil)

	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	nh.GetName()(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)
	r := dto.NameResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &r)
	assert.NoError(t, err)
	assert.Equal(t, "name-123", r.Name)
	nameService.AssertExpectations(t)

}

func NameHandler_GetName_NameGetError(t *testing.T, nameService *nameServiceMock, nh NameHandler) {

	nameService.On("Get", "555").Return("", errors.New("don't exists"))
	request, err := http.NewRequest(http.MethodGet, "/get/name?id=555", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	nh.GetName()(rr, request)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "don't exists\n", rr.Body.String())
	nameService.AssertExpectations(t)
}

func NameHandler_GetName_WithoutId(t *testing.T, nameService *nameServiceMock, nh NameHandler) {

	// without id
	nameService.On("Get", "555").Return("", errors.New("don't exists"))
	request, err := http.NewRequest(http.MethodGet, "/get/name", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	nh.GetName()(rr, request)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "bad request\n", rr.Body.String())
	nameService.AssertExpectations(t)
}

func NameHandler_SetName_Success(t *testing.T, nameService *nameServiceMock, nh NameHandler) {

	nameService.On("Set", dto.SetNameRequest{
		ID:   "id",
		Name: "name",
	}).Return(nil)
	request, err := http.NewRequest(http.MethodGet, "/set/name?id=id&name=name", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	nh.SetName()(rr, request)

	assert.Equal(t, http.StatusAccepted, rr.Code)
	assert.Equal(t, "done", rr.Body.String())
	nameService.AssertExpectations(t)

}

func NameHandler_SetName_WithoutIdOrName(t *testing.T, nameService *nameServiceMock, nh NameHandler) {

	request, err := http.NewRequest(http.MethodGet, "/set/name?id=id", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	nh.SetName()(rr, request)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "bad request\n", rr.Body.String())
	nameService.AssertExpectations(t)

	request, err = http.NewRequest(http.MethodGet, "/set/name?name=name", nil)
	assert.NoError(t, err)

	rr = httptest.NewRecorder()
	nh.SetName()(rr, request)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "bad request\n", rr.Body.String())
	nameService.AssertExpectations(t)

}

func NameHandler_SetName_NameSetError(t *testing.T, nameService *nameServiceMock, nh NameHandler) {

	nameService.On("Set", dto.SetNameRequest{
		ID:   "internal",
		Name: "internal",
	}).Return(errors.New("internal error"))
	request, err := http.NewRequest(http.MethodGet, "/set/name?id=internal&name=internal", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	nh.SetName()(rr, request)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "internal error\n", rr.Body.String())
	nameService.AssertExpectations(t)
}

func NameHandler_RemoveName_Success(t *testing.T, nameService *nameServiceMock, nh NameHandler) {

	nameService.On("Remove", "id").Return(nil)
	request, err := http.NewRequest(http.MethodGet, "/remove/name?id=id", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	nh.RemoveName()(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "done", rr.Body.String())
	nameService.AssertExpectations(t)

}

func NameHandler_RemoveName_WithoutId(t *testing.T, nameService *nameServiceMock, nh NameHandler) {

	request, err := http.NewRequest(http.MethodGet, "/remove/name", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	nh.RemoveName()(rr, request)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "bad request\n", rr.Body.String())
	nameService.AssertExpectations(t)
}

func NameHandler_RemoveName_NameRemoveError(t *testing.T, nameService *nameServiceMock, nh NameHandler) {

	nameService.On("Remove", "remid").Return(errors.New("internal error"))
	request, err := http.NewRequest(http.MethodGet, "/remove/name?id=remid", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	nh.RemoveName()(rr, request)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "internal error\n", rr.Body.String())
	nameService.AssertExpectations(t)
}
