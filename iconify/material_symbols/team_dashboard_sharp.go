package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TeamDashboardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21v-8.25H3V21h7Zm2 0h9v-8.25h-9V21ZM3 10.75h18V3H3v7.75Z"/>`),
		g.Group(children),
	)
}