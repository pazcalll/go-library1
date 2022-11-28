package models

import (
	"fmt"
	"io"
	"io/ioutil"
	"library/db"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Img    string `json:"img_base64"`
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

		bytes, err := ioutil.ReadFile("./" + book.Img)
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
		// c.RealIP()

		book.Img = base64Encoding

		arrObj = append(arrObj, book)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	defer rows.Close()
	return res, nil
}

func UploadBook(c echo.Context) (string, error) {
	name := c.FormValue("name")
	author := c.FormValue("author")
	//------------
	// Read files
	//------------

	file, err := c.FormFile("img_url")
	time_sec := time.Now().UnixNano()
	img_url_str := "images/book/" + strconv.Itoa(int(time_sec)) + "_" + file.Filename
	if err != nil {
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(img_url_str)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	sqlStatement := "INSERT INTO books(`name`, `author`, `img_url`) VALUES(?, ?, ?)"

	con := db.CreateCon()
	stmnt, _ := con.Prepare(sqlStatement)
	_, err = stmnt.Exec(name, author, img_url_str)

	if err != nil {
		panic(err)
	}

	return fmt.Sprint(http.StatusOK), err
}
