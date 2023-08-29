package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvGuideSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm5-5h2v-6h1.75L14 16h2l3-8h-2.5L15 12.5L13.5 8H5v2h3v6Z"/>`),
		g.Group(children),
	)
}