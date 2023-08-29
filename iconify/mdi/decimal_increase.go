package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecimalIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 17l-3 3v-2h-6v-2h6v-2l3 3M9 5a3 3 0 0 1 3 3v3a3 3 0 0 1-3 3a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3m0 2a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m7-2a3 3 0 0 1 3 3v3a3 3 0 0 1-3 3a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3m0 2a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1M4 12a1 1 0 0 1 1 1a1 1 0 0 1-1 1a1 1 0 0 1-1-1a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}