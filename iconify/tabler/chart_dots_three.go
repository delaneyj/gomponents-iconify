package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartDotsThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7a2 2 0 1 0 4 0a2 2 0 1 0-4 0m11 8a2 2 0 1 0 4 0a2 2 0 1 0-4 0m1-9a3 3 0 1 0 6 0a3 3 0 1 0-6 0M3 18a3 3 0 1 0 6 0a3 3 0 1 0-6 0m6-1l5-1.5m-7.5-7l7.81 5.37M7 7l8-1"/>`),
		g.Group(children),
	)
}