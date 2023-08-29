package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19.15l-2-2V10h-7.15l-2-2H22v11.15ZM14 6V4H6.85l-2-2H16v4h-2Zm6 11.15L12.85 10H20v7.15Zm.575 6.275L19.15 22H8V10.85l-4-4V14h2v2H2V4.85L.575 3.425L2 2l20 20l-1.425 1.425Z"/>`),
		g.Group(children),
	)
}