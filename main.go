package main

import (
	"WareHouse/DBConnection"
	"WareHouse/Warehouse"
	_ "github.com/lib/pq"
)

func main() {
	db := DBConnection.GetConnection()
	//w := Warehouse.NewUserRepository(db)
	//fmt.Println(w)
	Warehouse.Api(db)

}
