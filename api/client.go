package api

import (
	"examine/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)
var Clients []*Client
var Up = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
type Client struct {
	Id       string
	Status  byte//0为未准备 1为准备
	Username string
	Coon     *websocket.Conn
	Typ      byte //0为红方 1为黑方
}

//HandleNewConnection function creates a new client and stores it
func HandleNewConnection(c *gin.Context) {
	coon,err := Up.Upgrade(c.Writer,c.Request,nil)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK,gin.H{
			"status":10022,
			"info":"failed",
		})
		return
	}else {
		if len(Clients)>=2 {
			c.JSON(http.StatusOK,gin.H{
				"info":"房间人数已满",
			})
			return
		}
		log.Println("new connection")
		if len(Clients)==0 {
			newClient := Client{
				Id:       GenUserId(),
				Username: "",
				Coon:     coon,
				Status:   0,
				Typ:      0,
			}
			Clients =append(Clients,&newClient)
		}
		newClient:= Client{
			Id :      GenUserId(),
			Username: "",
			Coon:     coon,
			Status:   0,
			Typ:      1,
		}
		Clients =append(Clients,&newClient)
		defer coon.Close()

	}
}
//GenUserId function generates a random 10 character ID
func GenUserId() string{
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	rand.Seed(time.Now().UnixNano())

	b:= make([]rune,10)
	for i := range b{
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func ChangeStatus(c *gin.Context,client *Client)  {
	if client.Status==0 {
		client.Status=1
		c.JSON(200,gin.H{
			"status":"已准备",
		})
		return
	}else {
		client.Status = 0
		c.JSON(200,gin.H{
			"status":"未准备",
		})
		return
	}
}
func (client *Client)StartListening() {
	if len(Clients)!=2 {
		return
	}
	for _,j:=range Clients {
		if j.Status == 0{
			return
		}
	}
	for {
		_,content,err := client.Coon.ReadMessage()
		if err != nil {
			log.Println(err)
			client.Coon.Close()
			return
		}
		str := strings.Split(string(content)," ")
		if len(str) != 2{//第一个数字指定要走的棋子，第二个数字决定case
			fmt.Println("can not parse it")
			return
		}
		IChoose,ICase := str[0],str[1]
		choose,_:=strconv.Atoi(IChoose)
		cases,_:=strconv.Atoi(ICase)
		service.Operate(choose,client,cases)
	}
}
