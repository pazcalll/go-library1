package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func GetStock(c echo.Context) error {
	return c.String(http.StatusOK, "STOK")
}

func BorrowBook(c echo.Context) error {
	return c.String(http.StatusOK, "BORROW")
}

func ReturnBook(c echo.Context) error {
	return c.String(http.StatusOK, "RETURN")
}

func GetBorrowReport(c echo.Context) error {
	return c.String(http.StatusOK, "BORROW REPORT")
}

func GetReturnDetail(c echo.Context) error {
	return c.String(http.StatusOK, "RETURN DETAIL")
}

// MEMBER + USER
func BookAll(c echo.Context) error {
	return c.String(http.StatusOK, "MASTER BOOK")
}

func UserAll(c echo.Context) error {
	return c.String(http.StatusOK, "MASTER USER")
}
