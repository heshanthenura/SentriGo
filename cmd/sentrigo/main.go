package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/heshanthenura/SentriGo/internal/capture"
)

var start bool = false

func main() {
	// fmt.Println("🚀 Starting SentriGo IDS...")

	// r := gin.Default()
	// r.Static("/static", "./web/static")
	// r.LoadHTMLGlob("web/templates/*")

	// r.GET("/", handleIndex)
	// r.POST("/interface", handleInterface)
	// r.POST("/start", handleStart)

	// fmt.Println("🌐 WebUI running at http://localhost:8080")

	// if err := r.Run(":8080"); err != nil {
	// 	log.Fatal(err)
	// }

	capture.SniffInterface(`\Device\NPF_{BA91C28E-D8E6-4B6E-9193-050B6C4F44DF}`)

}

func handleIndex(c *gin.Context) {
	devices := capture.GetAvailableDevices()
	c.HTML(200, "index.html", gin.H{
		"devices": devices,
	})
}

func handleInterface(c *gin.Context) {
	iface := c.PostForm("iface")
	fmt.Println("🖧 Selected interface:", iface)
	c.String(200, "Received interface: "+iface)
}

func handleStart(c *gin.Context) {

}
