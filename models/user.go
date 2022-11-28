package models

import (
	"fmt"
	"io"
	"library/db"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Img_url string `json:"img_url"`
}

func UserAll() (Response, error) {
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

		arrObj = append(arrObj, user)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	defer rows.Close()
	return res, nil
}

func UploadUser(c echo.Context) (string, error) {
	name := c.FormValue("name")
	//------------
	// Read files
	//------------

	file, err := c.FormFile("img_url")
	time_sec := time.Now().UnixNano()
	name_str := "pics/" + strconv.Itoa(int(time_sec)) + "_" + file.Filename
	if err != nil {
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(name_str)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	sqlStatement := "INSERT INTO users(`name`, `img_url`) VALUES(?, ?)"

	con := db.CreateCon()
	stmnt, _ := con.Prepare(sqlStatement)
	_, err = stmnt.Exec(name, name_str)

	if err != nil {
		panic(err)
	}

	return fmt.Sprint(http.StatusOK), err
}
