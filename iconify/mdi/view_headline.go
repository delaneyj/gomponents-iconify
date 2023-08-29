package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewHeadline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5v2h17V5M4 11h17V9H4m0 10h17v-2H4m0-2h17v-2H4v2Z"/>`),
		g.Group(children),
	)
}