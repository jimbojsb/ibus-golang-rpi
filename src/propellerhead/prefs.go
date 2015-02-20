package propellerhead

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type preferences struct {
	airplay struct {
		speaker_name string
	}
	audio struct {
		source string
	}
	mpd struct {
		library_path string
		playlist_path string
	}
}



func CreatePreferencesFile() {
	Logger().Info("Creating default prefs.json")
	p := new(preferences)
	p.airplay.speaker_name = "propellerhead"
	p.audio.source = "airplay"
}

func Prefs() preferences {

	prefsFile := GetWorkingDir() + "/prefs.json"
	if _, err := os.Stat(prefsFile); err != nil {
		CreatePreferencesFile()
	}

	jsonString, _ := ioutil.ReadFile(prefsFile)

	var p preferences
	json.Unmarshal(jsonString, &p)
	return p
}

func GetWorkingDir() string {
	workingDir, _ := os.Getwd()
	return workingDir
}

func (p *preferences) Save() {
	jsonString, _ := json.MarshalIndent(p, "", "    ")
	ioutil.WriteFile(GetWorkingDir() + "/prefs.json", jsonString, 0664)
}
