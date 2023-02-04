package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/cat/add", AddCat)

	r.GET("/api/cats", GetAllCats)

	r.GET("/api/cat/:id", GetCat)

	r.DELETE("/api/cat/:id", DeleteCat)

	r.PUT("/api/cat/:id", EditCat)

	r.Run("0.0.0.0:8888")
}