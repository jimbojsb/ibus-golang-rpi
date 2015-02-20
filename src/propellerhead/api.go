package propellerhead

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"encoding/json"
)

func ServeApi() {

	r := mux.NewRouter()

	r.HandleFunc("/ibus/send", func(res http.ResponseWriter, req *http.Request) {
		Logger().Info("api: POST /ibus/send")
		pkt := new(IbusPacket)
		pkt.Src = req.FormValue("src")
		pkt.Dest = req.FormValue("dest")
		pkt.Message = strings.Split(req.FormValue("msg"), " ")
		pkt.CaclulateAndSaveChecksum()
		IbusDevices().SerialInterface.Write(pkt)

	}).Methods("POST")

	r.HandleFunc("/ibus/convert-string", func(res http.ResponseWriter, req *http.Request) {
		Logger().Info("api: POST /ibus/convert-string")
		res.Header().Set("Content-Type", "application/json")
		json, _ := json.Marshal(map[string]string{"hex": strings.Join(stringAsHexStringSlice(req.FormValue("text")), " ")})
		res.Write(json)
	}).Methods("POST")

	r.HandleFunc("/system/update", func(res http.ResponseWriter, req *http.Request) {
		Logger().Info("api: POST /system/update")

	}).Methods("POST")

	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			Logger().Info("api: GET /")
			res.Header().Set("Content-Type", "application/json")
			json, _ := json.Marshal("Welcome to Propellerhead")
			res.Write(json)
		}).Methods("GET")

	r.HandleFunc("/system/version", func(res http.ResponseWriter, req *http.Request) {
		Logger().Info("api: GET /system/version")
		res.Header().Set("Content-Type", "application/json")
		json, _ := json.Marshal(map[string]int{"version": PROPELLERHEAD_VERSION})
		res.Write(json)
	}).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":3281", nil)
}
