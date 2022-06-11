package main

import (
	"examine/api"
	"examine/dao"
)

func main() {
	api.InitEngine()
	dao.InitDB()
}
