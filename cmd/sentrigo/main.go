package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/heshanthenura/SentriGo/internal/capture"
)

var iface string

var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	fmt.Println("🚀 Starting SentriGo IDS...")

	r := gin.Default()
	r.Static("/assets", "./web/frontend/assets")
	r.NoRoute(func(c *gin.Context) {
		c.File("./web/frontend/index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		handleWebSocket(c.Writer, c.Request)
	})

	fmt.Println("🌐 WebUI running at http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	// capture.SniffInterface(`\Device\NPF_{BA91C28E-D8E6-4B6E-9193-050B6C4F44DF}`)

}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	clients[conn] = true
	go readMessages(conn)
}

func readMessages(conn *websocket.Conn) {
	defer func() {
		conn.Close()
		delete(clients, conn)
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received from client: %s\n", string(msg))

		var data map[string]interface{}
		err = json.Unmarshal(msg, &data)
		if err != nil {
			log.Println("Invalid JSON:", err)
			continue
		}
		action, _ := data["action"].(string)
		if action == "start-stop" {
			handleStartStop()
		} else if action == "get-iface" {
			devices := capture.GetAvailableDevices()

			var ifaceList []map[string]interface{}
			for _, d := range devices {
				ifaceList = append(ifaceList, map[string]interface{}{
					"name":        d.Name,
					"description": d.Description,
				})
			}

			data := map[string]interface{}{
				"type": "iface-list",
				"data": ifaceList,
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Println("Error converting iface data:", err)
				continue
			}

			err = conn.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				log.Println("WebSocket send error:", err)
			}
		} else if action == "toggle-iface" {
			iface = (data["iface"].(string))
		}
	}
}

func handleIndex(c *gin.Context) {
	devices := capture.GetAvailableDevices()
	c.HTML(200, "index.html", gin.H{
		"devices": devices,
	})
}

func handleStartStop() {
	log.Println(iface)
	// capture.ToggleSniffing(iface)
}
