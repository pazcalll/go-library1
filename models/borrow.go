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
	BorrowedAt int    `json:"borrowed_at"`
	ReturnedAt int    `json:"returned_at"`
}

var borrow Borrow

func BorrowBook(c echo.Context) (int, error) {
	book_id := c.FormValue("book_id")
	user_id := c.FormValue("user_id")

	sqlStatement := "SELECT * FROM borrows WHERE book_id = ? AND user_id = ? AND returned_at = null"

	con := db.CreateCon()
	err := con.QueryRow(sqlStatement, book_id, user_id).
		Scan(&borrow.Id, &borrow.UserId, &borrow.BookId, &borrow.BorrowedAt, &borrow.ReturnedAt)

	if err != nil {
		return DoBorrowBook(c, con)
	}

	defer con.Close()
	return 200, nil
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
