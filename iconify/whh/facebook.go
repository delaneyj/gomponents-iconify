package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Facebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 448q26 0 45 18.5t19 45.5v64q0 27-19 45.5T512 640H320v320q0 26-19 45t-45 19h-64q-27 0-45.5-19T128 960V640H64q-27 0-45.5-18.5T0 576v-64q0-27 18.5-45.5T64 448h64V256q0-106 75-181T384 0h128q27 0 45.5 18.5T576 64v64q0 27-18.5 45.5T512 192H384q-27 0-45.5 18.5T320 256v192h192z"/>`),
		g.Group(children),
	)
}