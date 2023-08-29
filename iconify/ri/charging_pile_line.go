package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChargingPileLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 19h1v2H2v-2h1V4a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v8h2a2 2 0 0 1 2 2v4a1 1 0 1 0 2 0v-7h-2a1 1 0 0 1-1-1V6.414l-1.657-1.657l1.414-1.414l4.95 4.95A.997.997 0 0 1 22 9v9a3 3 0 1 1-6 0v-4h-2v5Zm-9 0h7V5H5v14Zm4-8h3l-4 6v-4H5l4-6v4Z"/>`),
		g.Group(children),
	)
}