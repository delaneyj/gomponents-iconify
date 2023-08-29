package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaBBoxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 10.5c0 .8-.7 1.5-1.5 1.5c.8 0 1.5.7 1.5 1.5V15a2 2 0 0 1-2 2H9V7h4a2 2 0 0 1 2 2v1.5M13 15v-2h-2v2h2m0-4V9h-2v2h2M3 5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5m2 0v14h14V5H5Z"/>`),
		g.Group(children),
	)
}