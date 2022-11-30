package routes

import (
	// "net/http"

	"github.com/labstack/echo/v4"

	admin "library/controllers"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", admin.GetUser)                      // untuk cek apakah server berjalan atau tidak
	e.GET("/stock", admin.GetStock)                // mendapatkan stok semua buku, tapi akan mendapat stok dari sebuah buku jika diberi book_id (int) sebagai parameter
	e.POST("/borrow", admin.BorrowBook)            // meminjam buku dengan book_id (int) dan user_id (int) sebagai parameter
	e.POST("/upload/user", admin.UserUpload)       // menambah user dengan name (string) dan img (file: .png, .jpg, .jpeg) sebagai param
	e.POST("/upload/book", admin.BookUpload)       // menambahkan buku dengan name (string), author (string), img (file: .png, .jpeg) sebagai param
	e.PUT("/return", admin.ReturnBook)             // mengembalikan buku yang dipinjam dengan id (int) sebagai param
	e.GET("/borrow-report", admin.GetBorrowReport) // mendapatkan laporan rekap semua buku yang dipinjam
	e.GET("/return-detail", admin.GetReturnDetail) // mendapat informasi detail dari buku yang dipinjam dengan parameter borrow_id (int)
	e.GET("/book-all", admin.BookAll)              // mendapat semua data buku dari database ditambah foto
	e.GET("/user-all", admin.UserAll)              // mendatap semua data user dari database ditambah foto

	return e
}
