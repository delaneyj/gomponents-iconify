package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContainStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3h6v2H4v14h4v2H2V3m5 14v-2h2v2H7m4 0v-2h2v2h-2m4 0v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}