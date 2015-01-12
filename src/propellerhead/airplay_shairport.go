package propellerhead

import (
	"os/exec"
)

func RunShairport(c chan bool) {
	cmd := exec.Command("/usr/local/bin/shairport", "-M", GetWorkingDir()+"/shairport", "-D", GetWorkingDir()+"/shairport", "-a", Prefs().Airplay.SpeakerName)

	go func() {
		Logger().Info("starting shairport")
		cmd.Run()
		Logger().Info("stopped shairport")
	}()

	go func(quit chan bool) {
		<-quit
		Logger().Info("received kill signal for shairport")
		cmd.Process.Kill()
	}(c)
}
