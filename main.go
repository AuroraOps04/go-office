package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	_ = http.ListenAndServe(":4000", engine)
}
