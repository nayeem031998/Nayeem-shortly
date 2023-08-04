package main 

import (
	//"urlshortnerService/database"
	"urlshortnerService/database"
	"urlshortnerService/router"
	//"github.com/gin-gonic/gin"
	//"urlshortnerService/router"
	//"github.com/joho/godotenv"
)

// func init() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		panic(err)
// 	}
// 	database.ConnectToDB()
// }

func main() {
	router.ClientRoutes()
	database.ConnectToDB()

}
