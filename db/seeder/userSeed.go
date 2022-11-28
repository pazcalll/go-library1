package seeder

import "github.com/bxcodec/faker/v4"

func (s Seed) UserSeed() {
	for i := 0; i < 10; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO users(name) VALUES (?)`)
		// execute query
		_, err := stmt.Exec(faker.FirstName())
		if err != nil {
			panic(err)
		}
	}
}
