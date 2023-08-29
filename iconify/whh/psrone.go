package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Psrone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 768H192q-80 0-136-56T0 576V192q0-80 56-136T192 0h640q80 0 136 56t56 136v384q0 80-56 136t-136 56zM448 256q0-53-37.5-90.5T320 128h-96q-13 0-22.5 9.5T192 160v448q0 13 9.5 22.5T224 640t22.5-9.5T256 608V448h64q26 0 45 18.5t19 45.5v96q0 13 9.5 22.5T416 640t22.5-9.5T448 608v-96q0-58-44-96q44-38 44-96v-64zm352 320h-32V256q0-53-37.5-90.5T640 128q-14 0-23 9.5t-9 22.5t9 22.5t23 9.5q27 0 45.5 18.5T704 256v320h-64q-13 0-22.5 9.5T608 608t9.5 22.5T640 640h160q13 0 22.5-9.5T832 608t-9.5-22.5T800 576zM320 384h-64V192h64q26 0 45 18.5t19 45.5v64q0 26-19 45t-45 19z"/>`),
		g.Group(children),
	)
}