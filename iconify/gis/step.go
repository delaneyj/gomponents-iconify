package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Step(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M2.451 5A2.5 3 0 0 0 0 8v84a2.5 3 0 0 0 3.365 2.815l95-42a2.5 3 0 0 0 0-5.63l-95-42a2.5 3 0 0 0-.914-.184Z" color="currentColor"/>`),
		g.Group(children),
	)
}