package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Numbers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10V3L6 5m0 11a2 2 0 1 1 4 0c0 .591-.601 1.46-1 2l-3 3h4m5-7a2 2 0 1 0 2-2a2 2 0 1 0-2-2m-8.5 0h3"/>`),
		g.Group(children),
	)
}