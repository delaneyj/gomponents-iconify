package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeSendBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h14v14H2V2m20 6v14H8v-4h10V8h4M4 4v10h10V4H4Z"/>`),
		g.Group(children),
	)
}