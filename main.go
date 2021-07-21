package main

import (
	"flea-market/router"
	"net/http"

	_ "flea-market/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	router.LoadApiRouter(r)

	r.Run(":8080")
}
