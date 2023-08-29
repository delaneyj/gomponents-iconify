package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewSequential(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5h18v4H3V5m0 5h18v4H3v-4m0 5h18v4H3v-4Z"/>`),
		g.Group(children),
	)
}