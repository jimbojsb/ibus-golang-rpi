package propellerhead

import (
	"bufio"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type AirplayDacpRemote struct {
	dacp_id       string
	active_remote string
	host          string
	port          string
}

func (dacp *AirplayDacpRemote) Listen(quit chan bool) {

	fifo, _ := os.Open(GetWorkingDir() + "/shairport/dacp")

	go func() {
		Logger().Info("Started shairport dacp listener")
		scanner := bufio.NewScanner(fifo)
		var lines []string
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				for _, el := range lines {
					parts := strings.Split(el, "=")
					switch parts[0] {
					case "dacp_id":
						dacp.dacp_id = parts[1]
					case "active_remote":
						dacp.active_remote = parts[1]
					}
				}
				dacp.resolveMdns()
				lines = make([]string, 0)
			} else {
				lines = append(lines, line)
			}
		}
	}()
	<-quit
	fifo.Close()
	Logger().Info("Stopped shairport dacp listener")
}

func (dacp *AirplayDacpRemote) FastForward() {
	dacp.IssueCommand("beginff")
}

func (dacp *AirplayDacpRemote) Rewind() {
	dacp.IssueCommand("beginrew")
}

func (dacp *AirplayDacpRemote) ResumePlaying() {
	dacp.IssueCommand("playresume")
}

func (dacp *AirplayDacpRemote) Next() {
	dacp.IssueCommand("nextitem")
}

func (dacp *AirplayDacpRemote) Previous() {
	dacp.IssueCommand("previtem")
}

func (dacp *AirplayDacpRemote) Pause() {
	dacp.IssueCommand("pause")
}

func (dacp *AirplayDacpRemote) Play() {
	dacp.IssueCommand("play")
}

func (dacp *AirplayDacpRemote) Stop() {
	dacp.IssueCommand("stop")
}

func (dacp *AirplayDacpRemote) IssueCommand(command string) {
	url := "http://" + dacp.host + ":" + dacp.port + "/ctrl-int/1/" + command
	Logger().Debug("dacp: " + url)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Active-Remote", dacp.active_remote)
	http.DefaultClient.Do(request)
}

func (dacp *AirplayDacpRemote) resolveMdns() {

	var hostname string
	var port string
	if runtime.GOOS == "darwin" {
		dacpRecordName := "iTunes_Ctrl_" + dacp.dacp_id + "._dacp._tcp.local"
		dacpRecordCmd := exec.Command("/usr/local/bin/gtimeout", "1", dacpRecordName, "srv")
		Logger().Debug(dacpRecordCmd)
		dacpRecordOutputBytes, _ := dacpRecordCmd.Output()
		dacpRecordParts := strings.Split(string(dacpRecordOutputBytes), " ")
		dacp.port = dacpRecordParts[2]

		hostname := strings.TrimSpace(dacpRecordParts[3])
		hostnameRecordCmd := exec.Command("/usr/bin/dig", "+short", "@224.0.0.251", "-p", "5353", hostname, "A")
		Logger().Debug(hostnameRecordCmd)
		hostnameRecordOutputBytes, _ := hostnameRecordCmd.Output()
		hostnameRecordOutput := string(hostnameRecordOutputBytes)
		dacp.host = strings.TrimSpace(hostnameRecordOutput)
	} else {

	}

}
