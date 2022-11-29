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

type NullabledBorrow struct {
	ReturnedAt sql.NullString
	BorrowedAt sql.NullString
}

func BorrowBook(c echo.Context) (int, error) {
	var borrow Borrow
	var nullabledBorrow NullabledBorrow
	book_id, _ := strconv.Atoi(c.FormValue("book_id"))
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))

	sqlStatement := "SELECT * FROM borrows WHERE `book_id` = ? AND `user_id` = ? AND `returned_at` IS NULL LIMIT 1"

	con := db.CreateCon()
	err := con.QueryRow(sqlStatement, book_id, user_id).
		Scan(&borrow.Id, &borrow.UserId, &borrow.BookId, &nullabledBorrow.BorrowedAt, &nullabledBorrow.ReturnedAt)

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

func ReturnBook(c echo.Context) (int, error) {
	if c.FormValue("id") == "" {
		return http.StatusBadRequest, errors.New("ID harus diisi")
	}
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return http.StatusBadRequest, errors.New("ID harus angka")
	}

	sqlSelect := "SELECT * FROM borrows WHERE `id` = ? LIMIT 1"

	var borrow Borrow
	var nullabledBorrow NullabledBorrow
	con := db.CreateCon()
	err = con.QueryRow(sqlSelect, id).
		Scan(&borrow.Id, &borrow.UserId, &borrow.BookId, &nullabledBorrow.BorrowedAt, &nullabledBorrow.ReturnedAt)

	borrow.BorrowedAt = nullabledBorrow.BorrowedAt.String
	borrow.ReturnedAt = nullabledBorrow.ReturnedAt.String
	if err != nil {
		return http.StatusBadRequest, errors.New("Record peminjaman belum ada")
	}

	if nullabledBorrow.ReturnedAt.String != "" {
		return http.StatusBadRequest, errors.New("Buku sudah dikembalikan")
	}

	if borrow.ReturnedAt == "" {
		sqlUpdate := "UPDATE borrows SET `returned_at` = NOW() WHERE `id` = ?"
		stmntUpdate, _ := con.Prepare(sqlUpdate)
		_, err := stmntUpdate.Exec(id)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		sqlUpdate = "UPDATE stocks SET `stock` = CASE WHEN `stock`<`stock_max` THEN `stock`+1 ELSE `stock` END WHERE `book_id`=?"
		stmntUpdate1, _ := con.Prepare(sqlUpdate)
		_, err = stmntUpdate1.Exec(borrow.BookId)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	return http.StatusInternalServerError, nil
	// return http.StatusInternalServerError, c.JSON(500, borrow)
}
