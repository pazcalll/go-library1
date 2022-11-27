package seeder

import "github.com/bxcodec/faker/v4"

func (s Seed) BookSeed() {
	for i := 0; i < 100; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO books(name, author, img_url) VALUES (?,?,?)`)
		// execute query
		_, err := stmt.Exec(faker.Sentence(), faker.Name(), faker.URL())
		if err != nil {
			panic(err)
		}
	}
}
