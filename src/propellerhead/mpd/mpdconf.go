package mpd

import (
	"io/ioutil"
	"os"
	"propellerhead/prefs"
	"runtime"
)

func WriteMpdConf() {

	cwd := prefs.GetWorkingDir()
	mpdPath := cwd + "/mpd"

	var playlistDir, musicDir string

	if (runtime.GOOS == "darwin") {
		playlistDir = "/Volumes/PH_MUSIC/Playlists"
		musicDir = "/Volumes/PH_MUSIC/Music"
	} else if (runtime.GOOS == "linux") {
		playlistDir = "/music/Playlists"
		musicDir = "/music/Music"
	}

	mpdConf := "db_file \"" + mpdPath + "/mpd.db\"\n"
	mpdConf += "log_file \"" + mpdPath + "/mpd.log\"\n"
	mpdConf += "pid_file \"" + mpdPath + "/mpd.pid\"\n"
	mpdConf += "state_file \"" + mpdPath + "/mpd.state\"\n"
	mpdConf += "sticker_file \"" + mpdPath + "/sticker.sqlite\"\n"
	mpdConf += "music_directory \"" + musicDir + "\"\n"
	mpdConf += "playlist_directory \"" + playlistDir + "\"\n"
	mpdConf += "zeroconf_name \"Propellerhead\"\n"

	if runtime.GOOS == "darwin" {
		mpdConf += "audio_output {\n"
		mpdConf += "\ttype \"osx\"\n"
		mpdConf += "\tname \"osx\"\n"
		mpdConf += "\tmixer_type \"software\"\n"
		mpdConf += "}\n"
	} else if runtime.GOOS == "linux" {
		mpdConf += "audio_output {\n"
		mpdConf += "\ttype \"pulse\"\n"
		mpdConf += "\tname \"pulseaudio\"\n"
		mpdConf += "\tmixer_type \"software\"\n"
		mpdConf += "}\n"
	}

	ioutil.WriteFile(mpdPath+"/mpd.conf", []byte(mpdConf), os.FileMode(0664))
}
