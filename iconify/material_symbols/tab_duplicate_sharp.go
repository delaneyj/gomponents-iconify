package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDuplicateSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22H2v-2h2v2Zm-2-4v-2h2v2H2Zm0-4v-2h2v2H2Zm0-4V8h2v2H2Zm0-4V4h2v2H2Zm4 16v-2h2v2H6Zm4 0v-2h2v2h-2Zm4 0v-2h2v2h-2Zm4 0v-2h2v2h-2ZM6 18V2h16v16H6Zm7-10h7V4h-7v4Z"/>`),
		g.Group(children),
	)
}