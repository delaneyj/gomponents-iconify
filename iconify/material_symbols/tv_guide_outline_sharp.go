package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvGuideOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Zm3-3h2v-6h1.75L14 16h2l3-8h-2.5L15 12.5L13.5 8H5v2h3v6Z"/>`),
		g.Group(children),
	)
}