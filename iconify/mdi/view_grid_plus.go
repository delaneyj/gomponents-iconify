package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGridPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3v8h8V3h-8M3 21h8v-8H3v8M3 3v8h8V3H3m10 13h3v-3h2v3h3v2h-3v3h-2v-3h-3v-2Z"/>`),
		g.Group(children),
	)
}