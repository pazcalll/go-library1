package models

import (
	"library/db"
	"net/http"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Img    string `json:"img"`
}

func BookAll() (Response, error) {
	var book Book
	var arrObj []Book
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `books`"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.Img)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, book)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	defer rows.Close()
	return res, nil
}
