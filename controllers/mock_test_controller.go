package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

// TestController インターフェースのモック
type MockTestController struct {
	mock.Mock
}

func (m *MockTestController) TestHandler(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}
