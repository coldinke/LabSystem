package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"lab_backend/internal/mqtt_client"
	"lab_backend/internal/router"
)

func handleMQTTConnection(client mqtt.Client) {
	// 订阅主题
	mqtt_client.Sub(client, "data/pub")
}

func main() {
	mqttc, err := mqtt_client.Connect("ubuntu1.orb.local", 1883, "server", "1234")
	if err != nil {
		panic(err)
	}
	defer mqttc.Disconnect(250)
	
	go handleMQTTConnection(mqttc)

	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}
