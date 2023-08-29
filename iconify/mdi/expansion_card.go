package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpansionCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7v1.5h1V17h1.5V7H2m4 0v9h1v1h7v-1h8V7H6m11.5 2a2.5 2.5 0 0 1 2.5 2.5a2.5 2.5 0 0 1-2.5 2.5a2.5 2.5 0 0 1-2.5-2.5A2.5 2.5 0 0 1 17.5 9Z"/>`),
		g.Group(children),
	)
}