package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphonealt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 890v70q0 27-18.5 45.5t-45 18.5t-45.5-18.5t-19-45.5v-70Q183 867 91.5 760.5T0 512q0-27 18.5-45.5T64 448t45.5 18.5T128 512q0 106 75 181t181 75t181-75t75-181q0-27 19-45.5t45.5-18.5t45 18.5T768 512q0 142-91.5 248.5T448 890zm32-378h96q0 80-56 136t-135.5 56t-136-56T192 512h96q13 0 22.5-9.5T320 480t-9.5-22.5T288 448h-96v-64h96q13 0 22.5-9.5T320 352t-9.5-22.5T288 320h-96v-64h96q13 0 22.5-9.5T320 224t-9.5-22.5T288 192h-96q0-80 56.5-136t136-56T520 56t56 136h-96q-13 0-22.5 9.5T448 224t9.5 22.5T480 256h96v64h-96q-13 0-22.5 9.5T448 352t9.5 22.5T480 384h96v64h-96q-13 0-22.5 9.5T448 480t9.5 22.5T480 512z"/>`),
		g.Group(children),
	)
}