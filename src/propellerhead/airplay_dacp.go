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

	strrev := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	if (runtime.GOOS == "darwin") {
		dacpRecordName := "iTunes_Ctrl_" + dacp.dacp_id + "._dacp._tcp.local"
		dacpRecordCmd := exec.Command("/usr/local/bin/gtimeout", "1", "dns-sd", "-q", dacpRecordName, "srv")
		Logger().Debug(dacpRecordCmd)
		dacpRecordOutputBytes, _ := dacpRecordCmd.Output()
		dacpRecordLines := strings.Split(string(dacpRecordOutputBytes), "\n")
		dacpRecordLine := strrev(dacpRecordLines[3]);
		dacpRecordColumns := strings.Split(dacpRecordLine, " ");
		dacpRecordHostname := strrev(dacpRecordColumns[0]);
		dacp.port = strrev(dacpRecordColumns[1]);

		hostnameRecordCmd := exec.Command("/usr/local/bin/gtimeout", "1", "dns-sd", "-q", dacpRecordHostname);
		Logger().Debug(hostnameRecordCmd)
		hostnameRecordOutputBytes, _ := hostnameRecordCmd.Output();
		hostnameRecordLines := strings.Split(string(hostnameRecordOutputBytes), "\n");
		hostnameRecordLine := strrev(hostnameRecordLines[3]);
		hostnameRecordColumns := strings.Split(hostnameRecordLine, " ");
		dacp.host = strrev(hostnameRecordColumns[0]);
	} else {
		dacpRecordCmd := exec.Command("/usr/bin/avahi-browse", "--resolve", "-p", "-t", "_dacp._tcp")
		Logger().Debug(dacpRecordCmd)
		dacpRecordOutputBytes, _ := dacpRecordCmd.Output()
		dacpRecordLines := strings.Split(string(dacpRecordOutputBytes), "\n")
		dacpRecordLine := dacpRecordLines[1];
		dacpRecordColumns := strings.Split(dacpRecordLine, ";");
		dacp.host = dacpRecordColumns[7];
		dacp.port = dacpRecordColumns[6];
	}
	Logger().Info("DACP Discovered: " + dacp.host + ":" + dacp.port);

}
