package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RMobiledataSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 10V2h7v5.2H7.8L9 10H7L5.85 7.35H4V10H2Zm2-4.65h3V4H4v1.35Z"/>`),
		g.Group(children),
	)
}