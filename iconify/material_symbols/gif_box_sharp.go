package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GifBoxSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm4.5-7h3v-2h-1v1h-1v-2h2v-1h-3v4Zm4 0h1v-4h-1v4Zm2 0h1v-1.5H16v-1h-1.5V11h2v-1h-3v4Z"/>`),
		g.Group(children),
	)
}