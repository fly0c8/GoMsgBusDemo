package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"skidata.com/lib/libmsgbus/c_libmsgbus"
)

const (
	tenantId = "arniTheTenant"
	deviceId = "ltanar"
)

func msgHandler(msgId, msgType, range_first, range_last, createTs int64, topic, originTenantId, originBlName, originDeviceId, payload string) {
	fmt.Println("onMsg: ", topic, originTenantId, originBlName, originDeviceId, payload)
}
func networkEventHandler(eventCode int64, eventText, tenantId, nodeName string) {
	fmt.Println("onNetworkEvent: ", nodeName, eventText, tenantId)
}
func main() {

	if len(os.Args) != 2 {
		log.Fatal("Usage: ", os.Args[0], "<blname>")
	}
	blname := os.Args[1]

	if err := c_libmsgbus.Init(
		tenantId, blname, deviceId, "./data",
		msgHandler, networkEventHandler); err != nil {
		log.Fatal(err)
	}
	defer c_libmsgbus.Destroy()

	if err := c_libmsgbus.LoadNetworkCfgFromFile("device.cfg"); err != nil {
		log.Fatal(err)
	}

	if err := c_libmsgbus.Subscribe("testtopic", "", "1", c_libmsgbus.FLAG_SUBSCRIBE_NEWORIGINFROMSTART); err != nil {
		log.Fatal(err)
	}

	rdr := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter Message:")
		text, err := rdr.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Sending:" + text)
		if err := c_libmsgbus.Send("testtopic", text, 0, c_libmsgbus.AUTO_MSG_ID); err != nil {
			log.Fatal(err)
		}
	}
}
