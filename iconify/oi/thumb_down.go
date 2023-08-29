package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v4h1V0H0zm2 0v4c.28 0 .53.09.72.28c.19.19 1.15 2.12 1.28 2.38c.13.26.39.39.66.31c.26-.08.4-.36.31-.63c-.08-.26-.47-1.59-.47-1.84S4.72 4 5 4h1.5c.28 0 .5-.22.5-.5S5.97.31 5.97.31A.518.518 0 0 0 5.5 0H2z"/>`),
		g.Group(children),
	)
}