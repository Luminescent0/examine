package api

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

var addr = flag.String("addr", ":8080", "http service address")
var r = gin.Default()

func InitEngine() {
	r.POST("/login", login)
	r.GET("/register", register)
	r.GET("/joinRoom", JwtAuthMiddleware)
	r.GET("/ws", JwtAuthMiddleware, HandleNewConnection)
	err := r.Run(*addr)
	if err != nil {
		log.Fatal(err)
	}
}
