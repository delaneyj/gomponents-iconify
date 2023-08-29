package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Textcursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 128h-64q-24 0-44 47.5T320 256v512q0 33 20 80.5t44 47.5h64q26 0 45 19t19 45.5t-19 45t-45 18.5h-96q-29 0-62.5-41.5T256 896q0 45-33.5 86.5T160 1024H64q-27 0-45.5-18.5T0 960.5T18.5 915T64 896h64q24 0 44-47.5t20-80.5V256q0-33-20-80.5T128 128H64q-27 0-45.5-18.5T0 64.5T18.5 19T64 0h96q27 0 61.5 41.5T256 128q0-45 34.5-86.5T352 0h96q26 0 45 19t19 45.5t-19 45t-45 18.5z"/>`),
		g.Group(children),
	)
}