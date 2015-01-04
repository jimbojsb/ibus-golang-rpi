package propellerhead

import (
	"github.com/chuckpreslar/emission"
)

var eventEmitter *emission.Emitter = nil

func Emitter() (*emission.Emitter) {
	if (eventEmitter == nil) {
		eventEmitter = emission.NewEmitter()
	}
	return eventEmitter
}
