package prefs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Preferences struct {
	Airplay struct {
		SpeakerName string
	}
	Wireless struct {
		Ssid     string
		Password string
	}
	Ibus struct {
		Display string
	}
	State struct {
		AudioSource string
	}
}

func CreateDefault() {
	fmt.Println("Creating default prefs.json")
	p := new(Preferences)
	p.Airplay.SpeakerName = "propellerhead"
	p.Wireless.Ssid = "propellerhead"
	p.Wireless.Password = "propellerhead"
	p.Ibus.Display = "16x9"
	p.State.AudioSource = "airplay"
	Save(p)
}

func Get() Preferences {

	prefsFile := GetWorkingDir() + "/prefs.json"
	if _, err := os.Stat(prefsFile); err != nil {
		CreateDefault()
	}

	jsonString, err := ioutil.ReadFile(prefsFile)
	if err != nil {
		fmt.Println(err.Error)
	}

	var p Preferences
	json.Unmarshal(jsonString, &p)
	return p

}

func GetWorkingDir() string {
	workingDir, _ := os.Getwd()
	return workingDir
}

func Save(pref *Preferences) {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	jsonString, err := json.MarshalIndent(pref, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(workingDir+"/prefs.json", jsonString, 0664)
}
