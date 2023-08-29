package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LooksFourOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Zm8-2h2V7h-2v4h-2V7H9v6h4v4Z"/>`),
		g.Group(children),
	)
}