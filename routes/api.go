package routes

import (
	// "net/http"

	"github.com/labstack/echo/v4"

	"library/controllers"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", controllers.GetUser)
	e.GET("/stock", controllers.GetStock)
	e.POST("/borrow", controllers.BorrowBook)
	e.PUT("/return", controllers.GetBorrowReport)
	e.GET("/borrow-report", controllers.GetBorrowReport)
	e.GET("/return-detail", controllers.GetReturnDetail)
	e.GET("/book-all", controllers.BookAll)
	e.GET("/user-all", controllers.UserAll)

	return e
}
