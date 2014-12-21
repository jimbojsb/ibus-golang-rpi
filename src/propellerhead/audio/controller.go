package audio

import (
	mpdclient "github.com/fhs/gompd/mpd"
	"propellerhead/airplay"
	"propellerhead/mpd"
	"strconv"
	"propellerhead/prefs"
)

const SOURCE_MPD = "mpd"
const SOURCE_AIRPLAY = "airplay"

type Status struct {
	Source   string
	Song     string
	Artist   string
	Album    string
	Duration int
	Position int
}

type Controller struct {
	source                 string
	airplayMetadata        airplay.MetadataListener
	airplayRemote          airplay.DacpRemote
	processControlChannels []chan bool
}

func (ac *Controller) SetSource(source string) {
	if ac.source != "" {
		ac.Stop()
	}

	for _, pcn := range ac.processControlChannels {
		pcn <- true
	}
	ac.processControlChannels = make([]chan bool, 0)

	ac.source = source

	if ac.source == SOURCE_MPD {
		mpdControlChannel := make(chan bool)
		go mpd.RunMpd(mpdControlChannel)
		ac.processControlChannels = append(ac.processControlChannels, mpdControlChannel)
	} else if ac.source == SOURCE_AIRPLAY {
		shairportControlChannel := make(chan bool)
		metadataControlChannel := make(chan bool)
		dacpControlChannel := make(chan bool)

		go airplay.RunShairport(shairportControlChannel)
		go ac.airplayMetadata.Listen(metadataControlChannel)
		go ac.airplayRemote.Listen(dacpControlChannel)

		ac.processControlChannels = append(ac.processControlChannels, shairportControlChannel, metadataControlChannel, dacpControlChannel)
	}

	p := prefs.Get()
	p.State.AudioSource = ac.source
	prefs.Save(&p)

	ac.Play()
}

func (ac *Controller) Stop() {
	if ac.source == SOURCE_MPD {
		client := getMpdClient()
		client.Stop()
	} else if ac.source == SOURCE_AIRPLAY {
		ac.airplayRemote.Stop()
	}
}

func (ac *Controller) Next() {
	if ac.source == SOURCE_MPD {
		client := getMpdClient()
		client.Next()
	} else if ac.source == SOURCE_AIRPLAY {
		ac.airplayRemote.Next()
	}
}

func (ac *Controller) Prev() {
	if ac.source == SOURCE_MPD {
		// mimic mpc cdprev functionality
		client := getMpdClient()
		status, _ := client.Status()
		elapsedTime, _ := strconv.ParseFloat(status["elapsed"], 32)
		if elapsedTime <= 3 {
			client.Previous()
		} else {
			songId, _ := strconv.Atoi(status["songid"])
			client.SeekId(songId, 0)
		}
	} else if ac.source == SOURCE_AIRPLAY {
		ac.airplayRemote.Previous()
	}
}

func (ac *Controller) Pause() {
	if ac.source == SOURCE_MPD {
		client := getMpdClient()
		client.Pause(true)
	} else if ac.source == SOURCE_AIRPLAY {
		ac.airplayRemote.Pause()
	}
}

func (ac *Controller) Play() {
	if ac.source == SOURCE_MPD {
	} else if ac.source == SOURCE_AIRPLAY {
		ac.airplayRemote.Play()
	}
}

//func (ac *AudioController) WaitAndDispatch(command chan string, result chan *AudioControllerStatus) {
//	for {
//		switch (<- command) {
//		case "nowplaying": result <- ac.getStatus()
//		case "quit": return
//		case "play": ac.play()
//		case "pause": ac.pause()
//		case "prev": ac.prev()
//		case "next": ac.next()
//		}
//	}
//}

func (ac *Controller) GetStatus() *Status {
	status := new(Status)
	if ac.source == SOURCE_MPD {
		client := getMpdClient()
		song, _ := client.CurrentSong()
		status.Artist = song["Artist"]
		status.Album = song["Album"]
		status.Song = song["Title"]
		duration, _ := strconv.Atoi(song["Time"])
		status.Duration = duration
	} else if ac.source == SOURCE_AIRPLAY {
		status.Song = ac.airplayMetadata.Title
		status.Artist = ac.airplayMetadata.Artist
		status.Album = ac.airplayMetadata.Album
	}
	status.Source = ac.source
	return status
}

func getMpdClient() *mpdclient.Client {
	client, _ := mpdclient.Dial("tcp", "localhost:6600")
	return client
}
