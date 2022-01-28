package main

import (
	"log"
	"os"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/mdp/qrterminal/v3"
)

func turnLightOn() {
	log.Println("Turn Light On")
}

func turnLightOff() {
	log.Println("Turn Light Off")
}

func main() {

	bridge := accessory.NewBridge(accessory.Info{
		Name: "HK Test Bridge",
		Manufacturer: "BetaLixT",
		ID: 1,
	})
	bulb := accessory.NewLightbulb(accessory.Info{
		Name:         "Test Light Bulb",
		Manufacturer: "BetaLixT",
		ID: 2,
	})

	bulb.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on {
			turnLightOn()
		} else {
			turnLightOff()
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: "32193243"}, bridge.Accessory, bulb.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	pUri, err := t.XHMURI()
	if err != nil {
		log.Fatal("Failed to generate xhmuri")
		log.Fatal(err)
	} else {
		qrterminal.Generate(pUri, qrterminal.L, os.Stdout)
	}

	t.Start()
}
