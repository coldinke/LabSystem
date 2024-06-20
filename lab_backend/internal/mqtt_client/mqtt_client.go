package mqtt_client

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"lab_backend/internal/models"
	"strconv"
)

type jsonData struct {
	Fire        int `json:"fire"`
	Smoke       int `json:"smoke"`
	Temperature int `json:"temperature"`
	Humidity    int `json:"humidity"`
	FireRelay   int `json:"fire_relay_status"`
	SmokeRelay  int `json:"smoke_relay_status"`
}

type mqttMessage struct {
	LabID uint     `json:"labID"`
	Data  jsonData `json:"data"`
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Topic: %s | %s\n", msg.Topic(), msg.Payload())

	// 解析 JSON 数据
	var mqttreq mqttMessage
	err := json.Unmarshal(msg.Payload(), &mqttreq)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
	}
	// 保存传感器数据
	sensorData := models.SensorData{
		Name:        "labSensor" + strconv.Itoa(int(mqttreq.LabID)),
		LabID:       mqttreq.LabID,
		Fire:        mqttreq.Data.Fire,
		Smoke:       mqttreq.Data.Smoke,
		Temperature: mqttreq.Data.Temperature,
		Humidity:    mqttreq.Data.Humidity,
	}
	// 写入数据库
	models.SaveSensorData(models.DB, sensorData)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %+v", err)
}

func Connect(broker string, port int, username string, password string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.SetClientID("go_mqtt_client")
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port)) // 设置连接的 broker
	opts.SetKeepAlive(60)
	opts.SetUsername(username)                       // 设置用户名
	opts.SetPassword(password)                       // 设置密码
	opts.SetDefaultPublishHandler(messagePubHandler) // 设置默认的消息处理函数
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}

func Sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Suscribed to LWT %s", topic)
}

func Publish(client mqtt.Client, topic, payload string) {
	token := client.Publish(topic, 0, false, payload)
	token.Wait()
	if token.Error() != nil {
		// 处理发布失败的情况
		fmt.Println("MQTT发布失败:", token.Error())
	}
}
