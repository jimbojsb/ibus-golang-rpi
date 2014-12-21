package airplay

import (
	"fmt"
	"os/exec"
	"propellerhead/prefs"
)

func RunShairport(c chan bool) {
	cmd := exec.Command(prefs.GetWorkingDir()+"/bin/shairport", "-M", prefs.GetWorkingDir()+"/shairport", "-D", prefs.GetWorkingDir()+"/shairport", "-a", prefs.Get().Airplay.SpeakerName)

	go func() {
		fmt.Println("started shairport")
		cmd.Run()
		fmt.Println("stopped shairport")
	}()

	go func(quit chan bool) {
		<-quit
		fmt.Println("Received kill signal for shairport")
		cmd.Process.Kill()
	}(c)
}
