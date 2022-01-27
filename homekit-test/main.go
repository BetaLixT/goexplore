package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"log"
)

func turnLightOn() {
	log.Println("Turn Light On")
}

func turnLightOff() {
	log.Println("Turn Light Off")
}

func main() {
	info := accessory.Info{
		Name:         "Test Light Bulb",
		Manufacturer: "BetaLixT",
	}

	acc := accessory.NewLightbulb(info)

	acc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on {
			turnLightOn()
		} else {
			turnLightOff()
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: "32191143"}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
