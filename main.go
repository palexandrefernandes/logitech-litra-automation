package main

import (
	"log"
	"time"

	"github.com/keybase/client/go/lsof"
)

const (
	CAMERA             = "/dev/video2"
	TICK_DURATION_SECS = 1
)

func CameraInUse() bool {
	// The error is not important, usually it just means that no
	// process is using the target camera.
	processes, _ := lsof.MountPoint(CAMERA)

	return len(processes) > 0
}

func SetupSimpleLogging() {
	log.SetPrefix("[Litra] ")
}

func main() {
	SetupSimpleLogging()

	log.Println("starting...")
	var device *Lightbar
	lightConnected := false
	isInternalStateActive := false
	tick := time.Tick(time.Second * TICK_DURATION_SECS)

	for range tick {
		var err error
		if !lightConnected {
			log.Println("trying to connect")
			device = ConnectToLightBar()
			isInternalStateActive = device.IsOn()
			lightConnected = true
		}

		if lightConnected {
			cameraActive := CameraInUse()
			if !isInternalStateActive && cameraActive {
				log.Println("turn on")
				err = device.TurnOn()
				isInternalStateActive = true
			} else if isInternalStateActive && !cameraActive {
				log.Println("turn off")
				err = device.TurnOff()
				isInternalStateActive = false
			}
		}

		if err != nil {
			log.Println("connection lost")
			device.Close()
			lightConnected = false
		}
	}
}
