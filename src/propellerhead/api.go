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

	http.Handle("/", r)
	http.ListenAndServe(":3281", nil)
}
