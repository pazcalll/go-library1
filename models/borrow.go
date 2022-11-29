package models

import (
	"database/sql"
	"errors"
	"library/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Borrow struct {
	Id         int    `json:"id"`
	UserId     string `json:"user_id"`
	BookId     int    `json:"book_id"`
	BorrowedAt string `json:"borrowed_at"`
	ReturnedAt string `json:"returned_at"`
}

type Nullabled struct {
	ReturnedAt sql.NullString
	BorrowedAt sql.NullString
}

func BorrowBook(c echo.Context) (int, error) {
	var borrow Borrow
	var nullabled Nullabled
	book_id, _ := strconv.Atoi(c.FormValue("book_id"))
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))

	sqlStatement := "SELECT * FROM borrows WHERE `book_id` = ? AND `user_id` = ? AND `returned_at` IS NULL LIMIT 1"

	con := db.CreateCon()
	err := con.QueryRow(sqlStatement, book_id, user_id).
		Scan(&borrow.Id, &borrow.UserId, &borrow.BookId, &nullabled.BorrowedAt, &nullabled.ReturnedAt)

	if err != nil {
		// return 400, err
		return DoBorrowBook(c, con)
	}

	// defer con.Close()
	return http.StatusBadRequest, errors.New("User belum mengembalikan buku ini")
	// return 400, c.JSON(http.StatusBadRequest, borrow)
}

func DoBorrowBook(c echo.Context, con *sql.DB) (int, error) {
	user_id_int, _ := strconv.Atoi(c.FormValue("user_id"))
	book_id_int, _ := strconv.Atoi(c.FormValue("book_id"))

	sqlInsert := "INSERT INTO borrows(`user_id`, `book_id`) VALUES(?, ?)"

	stmntInsert, _ := con.Prepare(sqlInsert)
	_, err := stmntInsert.Exec(user_id_int, book_id_int)

	if err != nil {
		return http.StatusBadRequest, errors.New("ID User atau ID Buku tidak terdaftar")
	}

	sqlUpdate := "UPDATE stocks set `stock` = CASE WHEN `stock` > 0 THEN `stock`-1 ELSE `stock` END WHERE `book_id` = ?"
	stmntUpdate, _ := con.Prepare(sqlUpdate)
	_, err = stmntUpdate.Exec(book_id_int)

	if err != nil {
		return 500, err
	}
	return 200, nil
}
