package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17h12l-3.75-5l-3 4L9 13l-3 4Zm-3 4V3h18v18H3Z"/>`),
		g.Group(children),
	)
}