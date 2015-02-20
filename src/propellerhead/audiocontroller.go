package propellerhead

import (
	mpdclient "github.com/fhs/gompd/mpd"
	"strconv"
)

const AUDIO_SOURCE_MPD = "mpd"
const AUDIO_SOURCE_AIRPLAY = "airplay"
const EVENT_AUDIO_SOURCE_CHANGED = "audio_source_change"

type AudioStatus struct {
	Source   string
	Song     string
	Artist   string
	Album    string
	Duration int
	Position int
}

type AudioController struct {
	source                 string
	airplayMetadata        AirplayMetadataListener
	airplayRemote          AirplayDacpRemote
	processControlChannels []chan bool
}

func NewAudioController() (*AudioController) {
	ac := new(AudioController)
	ac.SetSource(Prefs().audio.source)
	ac.bindEvents()
	return ac
}

func (ac *AudioController) bindEvents() {
	Emitter().On(EVENT_IBUS_MFW_NEXT_RELEASE, ac.Next)
	Emitter().On(EVENT_IBUS_MFW_PREV_PUSH, ac.Prev)
}

func (ac *AudioController) SetSource(source string) {

	Emitter().Emit(EVENT_AUDIO_SOURCE_CHANGED, source)

	ac.Stop()

	for _, pcn := range ac.processControlChannels {
		pcn <- true
	}
	ac.processControlChannels = make([]chan bool, 0)

	ac.source = source

	if ac.source == AUDIO_SOURCE_MPD {
		mpdControlChannel := make(chan bool)
		go RunMpd(mpdControlChannel)
		ac.processControlChannels = append(ac.processControlChannels, mpdControlChannel)
	} else if ac.source == AUDIO_SOURCE_AIRPLAY {
		shairportControlChannel := make(chan bool)
		metadataControlChannel := make(chan bool)
		dacpControlChannel := make(chan bool)

		go RunShairport(shairportControlChannel)
		go ac.airplayMetadata.Listen(metadataControlChannel)
		go ac.airplayRemote.Listen(dacpControlChannel)

		ac.processControlChannels = append(ac.processControlChannels, shairportControlChannel, metadataControlChannel, dacpControlChannel)
	}

	p := Prefs()
	p.audio.source = ac.source
	p.Save()
}

func (ac *AudioController) Stop() {
	if ac.source == AUDIO_SOURCE_MPD {
		client := getMpdClient()
		client.Stop()
	} else if ac.source == AUDIO_SOURCE_AIRPLAY {
		ac.airplayRemote.Stop()
	}
}

func (ac *AudioController) Next() {
	Logger().Info("audiocontroller next")
	if ac.source == AUDIO_SOURCE_MPD {
		client := getMpdClient()
		client.Next()
	} else if ac.source == AUDIO_SOURCE_AIRPLAY {
		ac.airplayRemote.Next()
	}
}

func (ac *AudioController) Prev() {
	if ac.source == AUDIO_SOURCE_MPD {
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
	} else if ac.source == AUDIO_SOURCE_AIRPLAY {
		ac.airplayRemote.Previous()
	}
}

func (ac *AudioController) Pause() {
	if ac.source == AUDIO_SOURCE_MPD {
		client := getMpdClient()
		client.Pause(true)
	} else if ac.source == AUDIO_SOURCE_AIRPLAY {
		ac.airplayRemote.Pause()
	}
}

func (ac *AudioController) Play() {
	if ac.source == AUDIO_SOURCE_MPD {
	} else if ac.source == AUDIO_SOURCE_AIRPLAY {
		ac.airplayRemote.Play()
	}
}

func (ac *AudioController) GetStatus() (*AudioStatus) {
	status := new(AudioStatus)
	if ac.source == AUDIO_SOURCE_MPD {
		client := getMpdClient()
		song, _ := client.CurrentSong()
		status.Artist = song["Artist"]
		status.Album = song["Album"]
		status.Song = song["Title"]
		duration, _ := strconv.Atoi(song["Time"])
		status.Duration = duration
	} else if ac.source == AUDIO_SOURCE_AIRPLAY {
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
