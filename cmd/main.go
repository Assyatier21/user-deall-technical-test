package main

import (
	"fmt"
	"user/config"
	cons "user/models"
	"user/routes"
)

func main() {
	config.InitDB()
	echo := routes.GetRoutes()

	addres := cons.Addres
	port := cons.Port

	host := fmt.Sprintf("%s:%s", addres, port)

	_ = echo.Start(host)
}
