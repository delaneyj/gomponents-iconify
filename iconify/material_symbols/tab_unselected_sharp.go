package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabUnselectedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20v-2h2v2h-2ZM8 6V4h2v2H8Zm4 4V4h10v6H12Zm6 10v-2h2v-2h2v4h-4Zm-8 0v-2h2v2h-2Zm10-6v-2h2v2h-2ZM2 16v-2h2v2H2Zm0-4v-2h2v2H2Zm0 8v-2h2v2H2ZM2 8V4h4v2H4v2H2Zm4 12v-2h2v2H6Z"/>`),
		g.Group(children),
	)
}