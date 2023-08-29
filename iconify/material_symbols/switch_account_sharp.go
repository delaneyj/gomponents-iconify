package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchAccountSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 11q1.25 0 2.125-.875T17 8q0-1.25-.875-2.125T14 5q-1.25 0-2.125.875T11 8q0 1.25.875 2.125T14 11Zm-8 7V2h16v16H6Zm-4 4V6h2v14h14v2H2Zm6-6h12q-1.05-1.425-2.638-2.212T14 13q-1.775 0-3.363.788T8 16Z"/>`),
		g.Group(children),
	)
}