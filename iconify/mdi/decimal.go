package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Decimal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 7a3 3 0 0 0-3 3v3a3 3 0 0 0 6 0v-3a3 3 0 0 0-3-3m1 6a1 1 0 0 1-2 0v-3a1 1 0 0 1 2 0m6-3a3 3 0 0 0-3 3v3a3 3 0 0 0 6 0v-3a3 3 0 0 0-3-3m1 6a1 1 0 0 1-2 0v-3a1 1 0 0 1 2 0M6 15a1 1 0 1 1-1-1a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}