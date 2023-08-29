package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptionOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20V4h18v16H3Zm2-2h14V6H5v12Zm1-3h5v-2H9.5v.5h-2v-3h2v.5H11V9H6v6Zm7 0h5v-2h-1.5v.5h-2v-3h2v.5H18V9h-5v6Zm-8 3V6v12Z"/>`),
		g.Group(children),
	)
}