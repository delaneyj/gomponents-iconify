package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeenhereOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 23l-8-6V2h16v15l-8 6Zm0-2.5l6-4.5V4H6v12l6 4.5ZM10.95 15l5.65-5.65l-1.4-1.45l-4.25 4.25l-2.1-2.1l-1.45 1.4L10.95 15ZM12 4H6h12h-6Z"/>`),
		g.Group(children),
	)
}