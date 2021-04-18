package service

import (
	"awesomeProject1/internal/handler/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type dataManagerMock struct {
	mock.Mock
}

func (m *dataManagerMock) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func (m *dataManagerMock) Set(key string, value interface{}) error {

	args := m.Called(key, value)
	return args.Error(0)

}

func (m *dataManagerMock) Remove(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func TestNameService(t *testing.T) {

	dm := new(dataManagerMock)
	ns := NameS{
		dataManager: dm,
	}

	NameService_Get_Success(t, dm, ns)
	NameService_Get_DataManagerError(t, dm, ns)

	NameService_Set_Success(t, dm, ns)
	NameService_Set_DataManagerError(t, dm, ns)

	NameService_Remove_Success(t, dm, ns)
	NameService_Remove_DataManagerError(t, dm, ns)

}

func NameService_Get_Success(t *testing.T, dm *dataManagerMock, ns NameS) {
	id := "123"
	name := "name-123"
	dm.On("Get", id).Return(name, nil)

	r, err := ns.Get(id)
	assert.NoError(t, err)

	assert.Equal(t, r, name)

	dm.AssertExpectations(t)

}

func NameService_Get_DataManagerError(t *testing.T, dm *dataManagerMock, ns NameS) {
	id := "1234"
	dm.On("Get", id).Return("", errors.New("get error"))

	_, err := ns.Get(id)
	assert.Error(t, err)

	assert.Equal(t, "while getting user name error: get error", err.Error())
	dm.AssertExpectations(t)

	id = "12345"
	dm.On("Get", id).Return("", errors.New("redis: nil"))

	_, err = ns.Get(id)
	assert.Error(t, err)

	assert.Equal(t, "name wasn't set", err.Error())
	dm.AssertExpectations(t)

}

func NameService_Set_Success(t *testing.T, dm *dataManagerMock, ns NameS) {
	id := "123"
	name := "name-123"
	n := dto.SetNameRequest{
		ID:   id,
		Name: name,
	}

	dm.On("Set", id, name).Return(nil)

	err := ns.Set(n)
	assert.NoError(t, err)

	dm.AssertExpectations(t)
}

func NameService_Set_DataManagerError(t *testing.T, dm *dataManagerMock, ns NameS) {
	id := "1234"
	name := "name-1234"
	n := dto.SetNameRequest{
		ID:   id,
		Name: name,
	}

	dm.On("Set", id, name).Return(errors.New("set error"))

	err := ns.Set(n)
	assert.Error(t, err)

	assert.Equal(t, "while setting user name error: set error", err.Error())
	dm.AssertExpectations(t)
}

func NameService_Remove_Success(t *testing.T, dm *dataManagerMock, ns NameS) {
	id := "rm"
	dm.On("Remove", id).Return(nil)

	err := ns.Remove(id)
	assert.NoError(t, err)

	dm.AssertExpectations(t)

}

func NameService_Remove_DataManagerError(t *testing.T, dm *dataManagerMock, ns NameS) {
	id := "remove error"
	dm.On("Remove", id).Return(errors.New("remove error"))

	err := ns.Remove(id)
	assert.Error(t, err)

	assert.Equal(t, "while removing user name error: remove error", err.Error())

	dm.AssertExpectations(t)

}
