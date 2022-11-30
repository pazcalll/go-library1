package models

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"library/db"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Borrow struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	BookId     int    `json:"book_id"`
	BorrowedAt string `json:"borrowed_at"`
	ReturnedAt string `json:"returned_at"`
}

type NullabledBorrow struct {
	ReturnedAt sql.NullString
	BorrowedAt sql.NullString
}

type ReturnDetail struct {
	Id         int    `json:"id"`
	BorrowedAt string `json:"borrowed_at"`
	ReturnedAt string `json:"returned_at"`
	User       struct {
		Id   int    `json:"user_id"`
		Name string `json:"user_name"`
		Img  string `json:"user_img_url"`
	} `json: user`
	Book struct {
		BookId int    `json:"book_id"`
		Name   string `json:"book_name"`
		Author string `json:"book_author"`
		Img    string `json:"book_img_url"`
	} `json: book`
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

func GetBorrowReport() (Response, error) {
	var borrow Borrow
	var arrBorrow []Borrow
	var nullabledBorrow NullabledBorrow
	var res Response

	sqlInsert := "SELECT * FROM borrows"

	con := db.CreateCon()
	rows, err := con.Query(sqlInsert)

	if err != nil {
		res.Message = err.Error()
		res.Status = http.StatusInternalServerError
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&borrow.Id, &borrow.UserId, &borrow.BookId, &borrow.BorrowedAt, &nullabledBorrow.ReturnedAt)
		if err != nil {
			res.Message = err.Error()
			res.Status = http.StatusInternalServerError
			return res, err
		}
		borrow.ReturnedAt = nullabledBorrow.ReturnedAt.String

		arrBorrow = append(arrBorrow, borrow)
	}

	res.Message = "Success"
	res.Status = http.StatusOK
	res.Data = arrBorrow

	return res, nil
}

func GetReturnDetail(c echo.Context) (Response, error) {
	var res Response
	if c.FormValue("borrow_id") == "" {
		res.Status = http.StatusBadRequest
		return res, errors.New("borrow_id tidak boleh kosong")
		// return res, errors.New("borrow_id tidak boleh kosong")
	}
	borrow_id, err := strconv.Atoi(c.FormValue("borrow_id"))
	if err != nil {
		return res, errors.New("borrow_id harus angka")
	}

	var borrow Borrow
	var nullabledBorrow NullabledBorrow

	sqlSelect := "SELECT borrow.`id`, borrow.`user_id`, borrow.`book_id`, borrow.`borrowed_at`, borrow.`returned_at` FROM borrows AS borrow " +
		"LEFT JOIN users AS user ON user.`id` = borrow.`user_id` " +
		"LEFT JOIN books AS book ON book.`id` = borrow.`book_id` " +
		"WHERE borrow.`id` = ? AND borrow.`returned_at` IS NOT NULL " +
		"LIMIT 1"

	con := db.CreateCon()
	err = con.QueryRow(sqlSelect, borrow_id).Scan(&borrow.Id, &borrow.UserId, &borrow.BookId, &borrow.BorrowedAt, &nullabledBorrow.ReturnedAt)

	if err != nil {
		res.Message = "Record tidak ada"
		res.Status = http.StatusInternalServerError
		return res, errors.New("Record tidak ada")
	}

	borrow.ReturnedAt = nullabledBorrow.ReturnedAt.String

	var returnDetail ReturnDetail
	sqlSelect = "SELECT `id`, `name`, `img_url` FROM users WHERE `id` = ? LIMIT 1"
	con = db.CreateCon()
	err = con.QueryRow(sqlSelect, borrow.UserId).Scan(&returnDetail.User.Id, &returnDetail.User.Name, &returnDetail.User.Img)

	if err != nil {
		res.Message = err.Error()
		res.Status = http.StatusInternalServerError
		return res, err
	}

	bytes, err := ioutil.ReadFile("./" + returnDetail.User.Img)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	returnDetail.User.Img = base64Encoding

	sqlSelect = "SELECT `id`, `name`, `author`, `img_url` FROM books WHERE `id` = ? LIMIT 1"
	con = db.CreateCon()
	err = con.QueryRow(sqlSelect, borrow.BookId).Scan(&returnDetail.Book.BookId, &returnDetail.Book.Name, &returnDetail.Book.Author, &returnDetail.Book.Img)

	if err != nil {
		res.Message = err.Error()
		res.Status = http.StatusInternalServerError
		return res, err
	}

	bytes, err = ioutil.ReadFile("./" + returnDetail.Book.Img)
	if err != nil {
		log.Fatal(err)
	}

	// Determine the content type of the image file
	mimeType = http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	returnDetail.Book.Img = base64Encoding
	returnDetail.Id = borrow.Id
	returnDetail.BorrowedAt = borrow.BorrowedAt
	returnDetail.ReturnedAt = borrow.ReturnedAt

	res.Message = "Success"
	res.Status = http.StatusOK
	res.Data = returnDetail

	return res, nil
}
