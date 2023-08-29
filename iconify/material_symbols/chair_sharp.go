package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-2H1v-9h4v5h14v-5h4v9h-3v2h-2v-2H6v2H4Zm3-8V8H4V3h16v5h-3v5H7Z"/>`),
		g.Group(children),
	)
}