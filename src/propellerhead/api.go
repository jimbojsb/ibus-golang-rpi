package propellerhead

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func ServeApi() {

	r := mux.NewRouter()

	r.HandleFunc("/ibus/send", func(res http.ResponseWriter, req *http.Request) {
		Logger().Info("api: POST /ibus/send")
		pkt := new(IbusPacket)
		pkt.Src = req.FormValue("src")
		pkt.Dest = req.FormValue("dest")
		messageString := req.FormValue("msg")
		pkt.Message = strings.Split(messageString, " ")
		pkt.CaclulateAndSaveChecksum()
		IbusDevices().SerialInterface.Write(pkt)

	}).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":3281", nil)
}
