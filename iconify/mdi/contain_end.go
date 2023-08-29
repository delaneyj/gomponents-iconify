package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContainEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17v-2h2v2H7m4 0v-2h2v2h-2m4 0v-2h2v2h-2m7-14v18h-6v-2h4V5h-4V3h6Z"/>`),
		g.Group(children),
	)
}