import paho.mqtt.client as mqtt
import RPi.GPIO as GPIO
import time
import board
import adafruit_dht
import json
import logging

# MQTT Server address, port, auth configurations.
MQTT_BROKER = "172.20.10.5"
MQTT_PORT = 1883
MQTT_USER = "lab01"
MQTT_PASSWORD = "123456"
MQTT_PUB_TOPIC = "data/pub"

# GPIO PIN - BCM
FIRE_PIN =  5
SMOKE_PIN = 6
dhtDevice = adafruit_dht.DHT11(board.D18)
SOMKE_RELAY_PIN = 23
FIRE_RELAY_PIN = 12
LED_PIN = 25


# Status
SMOKE_RELAY_STATUS = False
FIRE_RELAY_STATUS = False


# setting logging config
def setup_logging(log_file="./logs/example.log", log_file_mode="w",log_level=logging.DEBUG):
    """ Set default logging config """
    logging.basicConfig(filename=log_file, filemode=log_file_mode,level=log_level,
        format='%(asctime)s - %(levelname)s - %(message)s',
            datefmt='%m/%d/%Y %I:%M:%S %p')

def on_connect(client, userdata, flags, reason_code, properties):
    if reason_code == 0:
        logger.info(f"Connected to MQTT Broker with result code {reason_code}")
    else:
        logger.warning(f"Failed to connect, return code %d\n", reason_code)
    client.subscribe("contorl")

def init_mqtt():
    client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)
    client.username_pw_set(username=MQTT_USER, password=MQTT_PASSWORD)

    client.on_connect = on_connect
    client.enable_logger()
    return client

def init_sensor():
    GPIO.setmode(GPIO.BCM)
    GPIO.setup(FIRE_PIN, GPIO.IN, pull_up_down=GPIO.PUD_UP)
    GPIO.setup(SMOKE_PIN, GPIO.IN, pull_up_down=GPIO.PUD_UP)
    GPIO.setup(SOMKE_RELAY_PIN, GPIO.OUT, initial=GPIO.LOW)
    GPIO.setup(FIRE_RELAY_PIN, GPIO.OUT, initial=GPIO.LOW)

def read_sensor_data():
    try:
        fire_value = GPIO.input(FIRE_PIN)
        smoke_value = GPIO.input(SMOKE_PIN)
    except Exception as error:
        print(f"Error reading GPIO pins: {e}")
        return None

    try:
        temperature_c = dhtDevice.temperature
        humidity = dhtDevice.humidity
    except RuntimeError as error:
        print(error.args[0])
        time.sleep(2.0)
    except Exception as error:
        
        return None

    sensor_data = {
        "fire": fire_value,
        "smoke": smoke_value,
        "temperature": temperature_c,
        "humidity": humidity
    }

    return sensor_data

def handle_publish_data(data):
    print(data)
    publish_data = {
        "labID": 1,
        "data": data
    }
    json_data = json.dumps(publish_data)
    return json_data


def control_smoke_relay(smoke_value):
    global SMOKE_RELAY_STATUS
    if smoke_value == 0 and not SMOKE_RELAY_STATUS:  # Smoke detected and relay is off
        GPIO.output(SOMKE_RELAY_PIN, GPIO.HIGH)
        SMOKE_RELAY_STATUS = True
        logger.info("Relay turned ON due to smoke detection.")
    elif smoke_value == 1 and SMOKE_RELAY_STATUS:  # No smoke and relay is on
        GPIO.output(SOMKE_RELAY_PIN, GPIO.LOW)
        SMOKE_RELAY_STATUS = False
        logger.info("Relay turned OFF as smoke cleared.")
    return SMOKE_RELAY_STATUS

def control_fire_relay(fire_value):
    global FIRE_RELAY_STATUS
    if fire_value == 0 and not FIRE_RELAY_STATUS:
        GPIO.output(FIRE_RELAY_PIN, GPIO.HIGH)
        FIRE_RELAY_STATUS = True
        logger.info("Relay turned ON due to fire detection.")
    elif fire_value == 1 and FIRE_RELAY_STATUS:  
        GPIO.output(FIRE_RELAY_PIN, GPIO.LOW)
        FIRE_RELAY_STATUS = False
        logger.info("Relay turned OFF as fire cleared.")
    return FIRE_RELAY_STATUS



setup_logging("./lab.log")
logger = logging.getLogger(__name__) 

def main():
    init_sensor()
    mqtt_client = init_mqtt()
    mqtt_client.connect(MQTT_BROKER, MQTT_PORT)
    try:
        # 保持MQTT客户端连接
        mqtt_client.loop_start()
        times = 0
        # 循环读取GPIO状态并发布MQTT消息
        while True:
            data = read_sensor_data()
            if data == None:
                continue
            fire_realy_status = control_fire_relay(data["fire"])
            data["fire_relay_status"] = 0 if fire_realy_status else 1
            smoke_relay_status = control_smoke_relay(data["smoke"])
            data["smoke_relay_status"] = 0 if smoke_relay_status else 1
            json_data = handle_publish_data(data)
            msg_info = mqtt_client.publish(MQTT_PUB_TOPIC, json_data)
            times += 1
            print(times)
            msg_info.wait_for_publish()
            time.sleep(10)

    except KeyboardInterrupt:
        print("程序被中断")

    finally:
        GPIO.cleanup()
        dhtDevice.exit()       
        mqtt_client.disconnect()
        mqtt_client.loop_stop()

if __name__ == "__main__":
    main()