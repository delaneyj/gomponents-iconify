package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roadsignright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1009.31 355l-206 142q-4 4-13 7q-7 8-22 8h-192v448q0 26-19 45t-45.5 19t-45-19t-18.5-45V512h-384q-27 0-45.5-18.5T.31 448V192q0-26 18.5-45t45.5-19h384V64q0-26 18.5-45t45-19t45.5 19t19 45v64h141q19-1 37 0h14q14 0 21 7q10 3 14 7l206 142q15 15 15 35.5t-15 35.5z"/>`),
		g.Group(children),
	)
}