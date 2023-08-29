package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewAgendaOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-8h18v8H3Zm2-2h14v-4H5v4Zm-2-8V3h18v8H3Zm2-2h14V5H5v4Zm0 6v4v-4ZM5 5v4v-4Z"/>`),
		g.Group(children),
	)
}