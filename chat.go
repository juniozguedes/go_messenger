package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
) 

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home Page")
	
}

func reader(conn *websocket.Conn){
	for {
		
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err) 
			return
		}

		conn.WriteMessage(messageType, p)
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return 
		}
	}
}

//Allow any con to websocket
func wsEndpoint(w http.ResponseWriter, r *http.Request){
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
	}

	//log.Println("Client Succesfully Connected... (Bateu no wsEndpoint do Back")
	log.printIn("alskdlasd")
	reader(ws)
}

func setupRoutes(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Go Websockets, começando na MAIN")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}