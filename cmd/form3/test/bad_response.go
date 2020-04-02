package test

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockReadCloser struct {
	mock.Mock
}

func (m *MockReadCloser) Read(p []byte) (int, error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *MockReadCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}

func BadResponse() *http.Response {
	mockReadCloser := MockReadCloser{}
	mockReadCloser.On("Read", mock.AnythingOfType("[]uint8")).Return(0, fmt.Errorf("error"))
	mockReadCloser.On("Close").Return(fmt.Errorf("error"))

	return &http.Response{
		Body: &mockReadCloser,
	}
}
