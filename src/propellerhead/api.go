package propellerhead

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"propellerhead/ibus"
	"propellerhead/audio"
	"propellerhead/prefs"
	"fmt"
)

type App struct {
	ibusOut chan *ibus.Packet
	ac *audio.Controller
}

func New(out chan *ibus.Packet, ac *audio.Controller) (*App) {
	app := new(App)
	app.ibusOut = out
	app.ac = ac
	return app
}

func (app *App) Serve() {

	r := mux.NewRouter()

	r.HandleFunc("/music/now-playing", func(res http.ResponseWriter, req *http.Request) {
			status := app.ac.GetStatus()
			res.Header().Set("Content-Type", "application/json")
			json, _ := json.Marshal(status)
			res.Write(json)
		}).Methods("GET")

	r.HandleFunc("/music/pause", func(res http.ResponseWriter, req *http.Request) {
			app.ac.Pause()
		}).Methods("POST")

	r.HandleFunc("/music/play", func(res http.ResponseWriter, req *http.Request) {
			app.ac.Play()
		}).Methods("POST")

	r.HandleFunc("/music/next", func(res http.ResponseWriter, req *http.Request) {
			app.ac.Next()
		}).Methods("POST")

	r.HandleFunc("/music/prev", func(res http.ResponseWriter, req *http.Request) {
			app.ac.Prev()
		}).Methods("POST")

	r.HandleFunc("/music/set-source", func(res http.ResponseWriter, req *http.Request) {
			app.ac.SetSource(req.FormValue("source"))
		}).Methods("POST")

	r.HandleFunc("/settings/ssid", func(res http.ResponseWriter, req *http.Request) {

		}).Methods("POST")

	r.HandleFunc("/settings/name", func(res http.ResponseWriter, req *http.Request) {

		}).Methods("POST")

	r.HandleFunc("/settings/name", func(res http.ResponseWriter, req *http.Request) {

		}).Methods("GET")

	r.HandleFunc("/ibus/send", func(res http.ResponseWriter, req *http.Request) {
			pkt := new(ibus.Packet)
			pkt.Src = req.FormValue("src")
			pkt.Dest = req.FormValue("dest")
			messageString := req.FormValue("msg")
			for _, el := range strings.Split(messageString, " ") {
				pkt.Message = append(pkt.Message, el)
			}
			pkt.CaclulateAndSaveChecksum()
			fmt.Printf("%+v", pkt)
			app.ibusOut <- pkt

		}).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(prefs.GetWorkingDir() + "/www")))
	http.Handle("/", r)
	http.ListenAndServe(":9000", nil)
}
