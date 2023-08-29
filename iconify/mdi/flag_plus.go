package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 14h2v3h3v2h-3v3h-2v-3h-3v-2h3v-3m-4.6-9H18v7c-2.22 0-4.16 1.21-5.2 3H11l-.4-2H5v7H3V3h9l.4 2Z"/>`),
		g.Group(children),
	)
}