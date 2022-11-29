package seeder

import "github.com/bxcodec/faker/v4"

func (s Seed) UserSeed() {
	for i := 0; i < 10; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO users(name, img_url) VALUES (?, ?)`)
		// execute query
		_, err := stmt.Exec(faker.FirstName(), "images/user/1669623322181445800_148764.png")
		if err != nil {
			panic(err)
		}
	}
}
