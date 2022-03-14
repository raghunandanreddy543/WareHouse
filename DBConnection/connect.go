package DBConnection

import (
	config "WareHouse/Config"
	"database/sql"
	"fmt"
	"log"
)

func GetConnection() *sql.DB {
	appConfig := config.New("./config.env")
	x := appConfig.DBConfig
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", x.DbUser, x.DbName, x.DbPassword, x.DbHost)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Couldn't Prepare the In Memory Connection")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Couldn't Connect DB")
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	var (
		Id       int
		Location string
		Name     string
		Quantity float32
		Unit     string
	)
	rows, err := db.Query("select * from warehouse;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Id, &Location, &Name, &Quantity, &Unit)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(Location, Name, Quantity, Unit)
	}
	return db
}
