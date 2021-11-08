package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Fajar3108/inventoryservice/database"
	"github.com/Fajar3108/inventoryservice/product"
	"github.com/Fajar3108/inventoryservice/receipt"
	_ "github.com/go-sql-driver/mysql"
)

const basePath = "/api"

func main(){
	database.SetupDatabase()
	product.SetupRoutes(basePath)
	receipt.SetupRoutes(basePath)
	fmt.Println("Starting web server at http://localhost:5000/")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}	
}
