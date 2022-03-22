package Warehouse

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func Api(db *sql.DB) {
	router := gin.Default()

	router.GET("/delhi/rice", func(context *gin.Context) {
		rows, err := db.Query("Select  PRODUCT_QUANTITY from warehouses Where product_name='Rice' and location='Delhi'")
		fmt.Println("Error is ", err)
		var (
			Quantity float32
		)
		for rows.Next() {
			err := rows.Scan(&Quantity)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(Quantity)
		}
		x := fmt.Sprintf("%s", Quantity)
		context.String(http.StatusOK, x)
	})
	router.POST("/delhi/rice", func(context *gin.Context) {
		query := fmt.Sprintf("insert into warehouses(LOCATION,PRODUCT_NAME,PRODUCT_QUANTITY,PRODUCT_UNIT) values(%s,%s,%s,%s)", "'Delhi'", "'Rice'", "200", "'KG'")
		fmt.Println(query)
		_, err := db.Query(query)
		fmt.Println("Error is ", err)
		if err != nil {
			panic("Error cannot update")
		}
		context.String(http.StatusCreated, "Ok")
	})

	router.Run(":8080")

}
