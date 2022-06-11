package api

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

var addr = flag.String("addr", ":8080", "http service address")

func InitEngine() {
	r := gin.Default()
	r.GET("/login", login)
	r.GET("/register", register)
	r.GET("/ws", JwtAuthMiddleware, HandleNewConnection)
	err := r.Run(*addr)
	if err != nil {
		log.Fatal(err)
	}
}
