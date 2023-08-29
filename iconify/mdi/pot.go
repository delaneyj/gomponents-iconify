package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19c0 1.1-.9 2-2 2H7c-1.1 0-2-.9-2-2v-7H3v-2h18v2h-2v7Z"/>`),
		g.Group(children),
	)
}