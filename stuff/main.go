package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"time"
)

func main() {
	drone := tello.NewDriver("8888")

	work := func() {
		drone.TakeOff()
		drone.SetVideoEncoderRate(3)
		drone.StartVideo()
		gobot.After(5*time.Second, func() {
			drone.Land()
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}
