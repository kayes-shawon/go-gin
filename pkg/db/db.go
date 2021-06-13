package db

import gopg "github.com/go-pg/pg/v10"

func ConnectDB()  *gopg.DB {
	db := gopg.Connect(&gopg.Options{
		Addr: "localhost:15432",
		User: "postgres",
		Password: "postgres",
		Database: "ginic",
	})
	return db
}
