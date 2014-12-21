package airplay

import (
	"bufio"
	"fmt"
	"os"
	"propellerhead/prefs"
	"strings"
)

type MetadataListener struct {
	Artist  string
	Title   string
	Album   string
	Artwork string
	Genre   string
}

func (mdl *MetadataListener) Listen(quit chan bool) {

	fifo, _ := os.Open(prefs.GetWorkingDir() + "/shairport/now_playing")

	go func() {
		fmt.Println("Started shairport metadata listener")
		scanner := bufio.NewScanner(fifo)
		var lines []string
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				for _, el := range lines {
					parts := strings.Split(el, "=")
					switch parts[0] {
					case "artist":
						mdl.Artist = parts[1]
					case "album":
						mdl.Album = parts[1]
					case "artwork":
						mdl.Artwork = parts[1]
					case "genre":
						mdl.Genre = parts[1]
					case "title":
						mdl.Title = parts[1]
					}
				}
				lines = make([]string, 0)
			} else {
				lines = append(lines, line)
			}
		}
	}()
	<-quit
	fifo.Close()
	fmt.Println("Stopped shairport metadata listener")
}
