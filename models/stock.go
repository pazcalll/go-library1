package models

import (
	"library/db"

	"github.com/labstack/echo/v4"
)

type Stock struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	StockMax int    `json:"stock_max"`
}

var res Response
var s Stock

func GetStock(c echo.Context) (Response, error) {
	var stocks []Stock

	if c.FormValue("book_id") != "" {
		con := db.CreateCon()
		sqlStatement := "SELECT s.id, b.name, s.stock, s.stock_max " +
			"FROM `stocks` as s LEFT JOIN `books` as b " +
			"ON b.id = s.id " +
			"WHERE s.id = ? " +
			"LIMIT 1"

		err := con.QueryRow(sqlStatement, c.FormValue("book_id")).Scan(&s.Id, &s.Name, &s.Stock, &s.StockMax)

		if err != nil {
			return res, err
		}

		res.Status = 200
		res.Message = "Success"
		res.Data = s

	} else {
		con := db.CreateCon()
		sqlStatement := "SELECT s.id, b.name, s.stock, s.stock_max " +
			"FROM `stocks` as s LEFT JOIN `books` as b " +
			"ON b.id = s.id"

		rows, err := con.Query(sqlStatement)

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&s.Id, &s.Name, &s.Stock, &s.StockMax)
			if err != nil {
				return res, err
			}
			stocks = append(stocks, s)
		}
		res.Status = 200
		res.Message = "Success"
		res.Data = stocks
	}

	return res, nil
}
