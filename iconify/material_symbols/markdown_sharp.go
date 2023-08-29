package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 15l3-3l-1.05-1.075l-1.2 1.2V9h-1.5v3.125l-1.2-1.2L13 12l3 3ZM2 20V4h20v16H2Zm3.5-5H7v-4.5h1v3h1.5v-3h1V15H12V9H5.5v6Z"/>`),
		g.Group(children),
	)
}