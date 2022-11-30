package admin

import (
	"library/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "server on")
}

func GetStock(c echo.Context) error {
	result, err := models.GetStock(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func BorrowBook(c echo.Context) error {
	var res models.Response
	status, err := models.BorrowBook(c)
	if err != nil {
		return c.JSON(status, map[string]interface{}{"status": status, "message": err.Error()})
	}
	res.Data = nil
	res.Status = 200
	res.Message = "Success"
	return c.JSON(http.StatusOK, res)
}

func ReturnBook(c echo.Context) error {
	var res models.Response
	status, err := models.ReturnBook(c)
	if err != nil {
		return c.JSON(status, map[string]interface{}{"status": status, "message": err.Error()})
	}
	res.Data = nil
	res.Status = 200
	res.Message = "Success"
	return c.JSON(http.StatusOK, res)
}

func GetBorrowReport(c echo.Context) error {
	result, err := models.GetBorrowReport()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": result.Status, "message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
	// return c.String(http.StatusOK, "BORROW REPORT")
}

func GetReturnDetail(c echo.Context) error {
	result, err := models.GetReturnDetail(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": result.Status, "message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

// MEMBER + BOOK
func BookAll(c echo.Context) error {
	result, err := models.BookAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func UserAll(c echo.Context) error {
	res, err := models.UserAll(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func BookUpload(c echo.Context) error {
	res, err := models.UploadBook(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func UserUpload(c echo.Context) error {
	res, err := models.UploadUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}
