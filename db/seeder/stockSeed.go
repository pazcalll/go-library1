package seeder

func (s Seed) StockSeed() {
	for i := 0; i < 100; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO stocks(book_id, stock, stock_max) VALUES (?, ?, ?)`)
		// execute query
		_, err := stmt.Exec(i+1, 10, 10)
		if err != nil {
			panic(err)
		}
	}
}
