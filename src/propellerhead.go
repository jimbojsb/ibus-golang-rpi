package main

import (
	"sync"
	"propellerhead/ibus"
	"propellerhead/prefs"
	"propellerhead/audio"
	"propellerhead/webapp"
	"os"
)

func main() {

	ttyPath := os.Args[1];

	wait := &sync.WaitGroup{}

	p := prefs.Get()

	ac := new(audio.Controller)
	ac.SetSource(p.State.AudioSource)

	wait.Add(1)
	ibusInterface := ibus.NewInterface(ac)
	go ibusInterface.Listen()

	wait.Add(1)
	app := webapp.New(ibusInterface.GetOutboundChannel(), ac)
	go app.Serve()

	wait.Wait()
}
