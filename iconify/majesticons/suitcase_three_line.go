package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 6V5a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v1m6 0h3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1M15 6H9m0 0H6a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h1m2-10v6m6 0v-6M7 20v1m0-1h10m0 0v1"/>`),
		g.Group(children),
	)
}