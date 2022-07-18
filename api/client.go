package api

import (
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
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Id       string
	Status   byte //0为未准备 1为准备
	Username string
	Coon     *websocket.Conn
	Typ      byte //0为红方 1为黑方
	Room     Room
}

//HandleNewConnection function creates a new client and stores it
func HandleNewConnection(c *gin.Context) {
	coon, err := Up.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": 10022,
			"info":   "failed",
		})
		return
	}
	log.Println("new connection")
	roomId := c.PostForm("roomId")
	newClient := Client{
		Id:       GenUserId(),
		Username: "",
		Coon:     coon,
		Status:   0,
		Typ:      0,
	}
	if roomId != "" {
		newClient.Room.Id = roomId
	} else {
		newClient.Room.Id = GenRandomNum()
	} //这里不对房间号是否存在作验证(我不会qaq),如果用户输的房间号是不存在的，就当执行了一次自定义的随机数吧(.
	Clients = append(newClient.Room.Clients, &newClient)
	var intChan chan int
	intChan = make(chan int, 1)
	intChan <- 1
	if len(newClient.Room.Clients) < 2 {
		<-intChan
	}
	if len(newClient.Room.Clients) == 2 {
		newClient.Room.Clients[0].Typ = 0
		newClient.Room.Clients[1].Typ = 1
		intChan <- 1
	}
	defer func(coon *websocket.Conn) {
		err := coon.Close()
		if err != nil {
			log.Println(err)
		}
	}(coon)

}

//GenUserId function generates a random 10 character ID
func GenUserId() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func GenRandomNum() string {
	var letters = []rune("1234567890")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func ChangeStatus(client *Client) {
	if client.Status == 0 {
		client.Status = 1
		return
	} else {
		client.Status = 0
		return
	}
}
func (client *Client) StartListening() {
	for {
		if len(client.Room.Clients) == 2 {
			client.Room.Clients[0].Typ = 0
			client.Room.Clients[1].Typ = 1
			break
		}
	}
	if Clients[0].Status != 1 || Clients[1].Status != 1 {
		for {
			_, content, err := client.Coon.ReadMessage()
			if err != nil {
				log.Println(err)
				err := client.Coon.Close()
				if err != nil {
					log.Println(err)
					return
				}
				return
			}
			if content != nil {
				ChangeStatus(client) //输入任意内容切换准备状态
			}
		}
	}
	InitBoard()
	for {
		Operate1(client)
	}
}
func read(client *Client) {
	_, content, err := client.Coon.ReadMessage()
	if err != nil {
		log.Println(err)
		err := client.Coon.Close()
		if err != nil {
			log.Println(err)
		}
		return
	}
	str := strings.Split(string(content), " ")
	if len(str) != 2 { //第一个数字指定要走的棋子，第二个数字决定case
		fmt.Println("can not parse it")
		return
	}
	IChoose, ICase := str[0], str[1]
	choose, _ := strconv.Atoi(IChoose)
	cases, _ := strconv.Atoi(ICase)
	Operate(choose, client, cases)
	err = client.Coon.WriteJSON(Board)
	if err != nil {
		log.Println(err)
	}
	return
}
func Operate1(client *Client) {
	var Chan chan byte
	Chan = make(chan byte, 1)
	arr := client.Room.Clients
	if arr[0].Typ == 1 {
		Chan <- 1
		read(arr[0])
	}
	read(arr[1])
	<-Chan
}
