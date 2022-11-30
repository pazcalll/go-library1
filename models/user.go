package models

import (
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"library/db"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Img_url string `json:"img_base64"`
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func UserAll(c echo.Context) (Response, error) {
	var user User
	var arrObj []User
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `users`"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Img_url)
		if err != nil {
			return res, err
		}

		bytes, err := ioutil.ReadFile("./" + user.Img_url)
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

		user.Img_url = base64Encoding
		arrObj = append(arrObj, user)
	}

	defer rows.Close()
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj
	return res, nil
}

func UploadUser(c echo.Context) (int, error) {
	name := c.FormValue("name")
	if name == "" {
		return http.StatusBadRequest, errors.New("Nama tidak boleh kosong")
	}
	//------------
	// Read files
	//------------

	file, err := c.FormFile("img")
	if err != nil {
		return http.StatusBadRequest, errors.New("File tidak boleh kosong")
	}

	if filepath.Ext(file.Filename) != ".jpeg" && filepath.Ext(file.Filename) != ".png" {
		return http.StatusBadRequest, errors.New("format file harus .png / .jpeg")
	}
	time_sec := time.Now().UnixNano()
	name_str := "images/user/" + strconv.Itoa(int(time_sec)) + "_" + file.Filename
	src, err := file.Open()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(name_str)
	if err != nil {
		return http.StatusInsufficientStorage, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return http.StatusInternalServerError, err
	}
	sqlStatement := "INSERT INTO users(`name`, `img_url`) VALUES(?, ?)"

	con := db.CreateCon()
	stmnt, _ := con.Prepare(sqlStatement)
	_, err = stmnt.Exec(name, name_str)

	if err != nil {
		panic(err)
	}

	return http.StatusOK, err
}
