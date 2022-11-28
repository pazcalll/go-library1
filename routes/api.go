package routes

import (
	// "net/http"

	"github.com/labstack/echo/v4"

	admin "library/controllers"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", admin.GetUser)
	e.GET("/stock", admin.GetStock)
	e.POST("/borrow", admin.BorrowBook)
	e.POST("/upload", admin.UserUpload)
	e.PUT("/return", admin.GetBorrowReport)
	e.GET("/borrow-report", admin.GetBorrowReport)
	e.GET("/return-detail", admin.GetReturnDetail)
	e.GET("/book-all", admin.BookAll)
	e.GET("/user-all", admin.UserAll)

	return e
}
