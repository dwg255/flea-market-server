package main

import (
	"flea-market/router"
	"log"
	"net/http"

	_ "flea-market/model"

	"github.com/gin-gonic/gin"
)

func main() {
	//fmt.Println("air works!")
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	router.LoadApiRouter(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
