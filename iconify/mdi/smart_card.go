package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3h20a2.07 2.07 0 0 1 2 2v14a2.07 2.07 0 0 1-2 2H2a2.07 2.07 0 0 1-2-2V5a2.07 2.07 0 0 1 2-2m6 10.91C6 13.91 2 15 2 17v1h12v-1c0-2-4-3.09-6-3.09M8 6a3 3 0 1 0 3 3a3 3 0 0 0-3-3m9 4v3h4v-3h-4"/>`),
		g.Group(children),
	)
}