package main

import (
	"log"
	"math"
	"time"

	"github.com/sstallion/go-hid"
)

const (
	LIGHTBAR_ID          = 0x046D
	VENDOR_ID            = 0xC901
	SERIAL               = "2231FE7015A8"
	COMMAND_DELAY_SECS   = 2.5
	RECONNECT_DELAY_SECS = 1
)

var IS_ON = []byte{0x11, 0xff, 0x04, 0x01}
var TURN_ON = []byte{0x11, 0xff, 0x04, 0x1c, 0x01}
var TURN_OFF = []byte{0x11, 0xff, 0x04, 0x1c, 0x00}

type Lightbar struct {
	hid *hid.Device
}

func (l Lightbar) IsOn() bool {
	_, err := l.hid.Write(IS_ON)
	if err != nil {
		return false
	}

	response := []byte{0x00, 0x00, 0x00, 0x00, 0x00}
	_, err = l.hid.Read(response)
	if err != nil {
		return false
	}

	return response[4] == 1
}

func (l Lightbar) sendCommand(command []byte) error {
	if _, err := l.hid.Write(command); err != nil {
		return err
	}

	time.Sleep(time.Duration(int64(math.Round(float64(time.Second.Microseconds()) * COMMAND_DELAY_SECS))))

	return nil
}

func (l Lightbar) Close() {
	l.hid.Close()
}

func (l Lightbar) TurnOn() error {
	return l.sendCommand(TURN_ON)
}

func (l Lightbar) TurnOff() error {
	return l.sendCommand(TURN_OFF)
}

func ConnectToLightBar() *Lightbar {
	var device *Lightbar
	for device == nil {
		deviceConnection, err := hid.Open(LIGHTBAR_ID, VENDOR_ID, SERIAL)

		if err == nil {
			log.Println("connected")
			device = &Lightbar{hid: deviceConnection}
		}

		time.Sleep(time.Second * RECONNECT_DELAY_SECS)
	}

	return device
}
