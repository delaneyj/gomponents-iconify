package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumericTenBoxMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5v16h16v2H3a2 2 0 0 1-2-2V5h2m13 8h2V7h-2v6m5-12H7c-1.1 0-2 .9-2 2v14a2 2 0 0 0 2 2h14c1.11 0 2-.89 2-2V3a2 2 0 0 0-2-2m-9 14h-2V7H8V5h4v10m8-2c0 1.11-.89 2-2 2h-2a2 2 0 0 1-2-2V7c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v6Z"/>`),
		g.Group(children),
	)
}