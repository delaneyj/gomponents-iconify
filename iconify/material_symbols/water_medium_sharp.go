package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterMediumSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.125 12.15q1.175-.575 2.438-.863T11.124 11q.75 0 1.488.1t1.462.3q1.25.35 1.913.475T17.4 12h.475l.875-8H5.25l.875 8.15ZM5.2 22L3 2h18l-2.2 20H5.2Z"/>`),
		g.Group(children),
	)
}