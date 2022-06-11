package api

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)
var addr = flag.String("addr",":8080","http service address")
func InitEngine()  {
	r := gin.Default()
	r.GET("/ws", HandleNewConnection)
	err := r.Run(*addr)
	if err != nil {
		log.Fatal(err)
	}
}