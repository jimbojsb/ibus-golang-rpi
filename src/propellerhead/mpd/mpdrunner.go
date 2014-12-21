package mpd

import (
	"fmt"
	"os/exec"
	"propellerhead/prefs"
)

func RunMpd(c chan bool) {
	WriteMpdConf()
	cmd := exec.Command("/usr/local/bin/mpd", "--no-daemon", prefs.GetWorkingDir()+"/mpd/mpd.conf")

	go func() {
		fmt.Println("started mpd")
		cmd.Run()
		fmt.Println("stopped mpd")
	}()

	go func(quit chan bool) {
		<-quit
		fmt.Println("Received kill signal for mpd")
		cmd.Process.Kill()
	}(c)
}
