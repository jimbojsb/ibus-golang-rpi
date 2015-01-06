package propellerhead

import (
	"fmt"
	"os/exec"
)

func RunShairport(c chan bool) {
	cmd := exec.Command(GetWorkingDir()+"/shairport", "-M", GetWorkingDir()+"/shairport", "-D", GetWorkingDir()+"/shairport", "-a", Prefs().Airplay.SpeakerName)

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
