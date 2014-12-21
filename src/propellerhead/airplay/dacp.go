package airplay

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"propellerhead/prefs"
	"strings"
)

type DacpRemote struct {
	dacp_id       string
	active_remote string
	host          string
	port          string
}

func (dacp *DacpRemote) Listen(quit chan bool) {

	fifo, _ := os.Open(prefs.GetWorkingDir() + "/shairport/dacp")

	go func() {
		fmt.Println("Started shairport dacp listener")
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
	fmt.Println("Stopped shairport dacp listener")
}

func (dacp *DacpRemote) FastForward() {
	dacp.IssueCommand("beginff")
}

func (dacp *DacpRemote) Rewind() {
	dacp.IssueCommand("beginrew")
}

func (dacp *DacpRemote) ResumePlaying() {
	dacp.IssueCommand("playresume")
}

func (dacp *DacpRemote) Next() {
	dacp.IssueCommand("nextitem")
}

func (dacp *DacpRemote) Previous() {
	dacp.IssueCommand("previtem")
}

func (dacp *DacpRemote) Pause() {
	dacp.IssueCommand("pause")
}

func (dacp *DacpRemote) Play() {
	dacp.IssueCommand("play")
}

func (dacp *DacpRemote) Stop() {
	dacp.IssueCommand("stop")
}

func (dacp *DacpRemote) IssueCommand(command string) {
	url := "http://" + dacp.host + ":" + dacp.port + "/ctrl-int/1/" + command
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Active-Remote", dacp.active_remote)
	http.DefaultClient.Do(request)
}

func (dacp *DacpRemote) resolveMdns() {

	dacpRecordName := "iTunes_Ctrl_" + dacp.dacp_id + "._dacp._tcp.local"
	dacpRecordCmd := exec.Command("/usr/bin/dig", "+short", "@224.0.0.251", "-p", "5353", dacpRecordName, "SRV")
	dacpRecordOutputBytes, _ := dacpRecordCmd.Output()
	dacpRecordParts := strings.Split(string(dacpRecordOutputBytes), " ")
	dacp.port = dacpRecordParts[2]

	hostname := strings.TrimSpace(dacpRecordParts[3])
	hostnameRecordCmd := exec.Command("/usr/bin/dig", "+short", "@224.0.0.251", "-p", "5353", hostname, "A")
	hostnameRecordOutputBytes, _ := hostnameRecordCmd.Output()
	hostnameRecordOutput := string(hostnameRecordOutputBytes)
	dacp.host = strings.TrimSpace(hostnameRecordOutput)
}
