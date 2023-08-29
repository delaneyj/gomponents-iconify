package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 2h14v1h1v1h1v14h-1v1h-1v1H4v-1H3v-1H2V4h1V3h1V2m0 14h1v-1h2v-1h8v1h2v1h1V5h-1V4H5v1H4v11m12 2v-1h-2v-1H8v1H6v1h10M9 5h4v1h1v1h1v4h-1v1h-1v1H9v-1H8v-1H7V7h1V6h1V5m3 3V7h-2v1H9v2h1v1h2v-1h1V8h-1Z"/>`),
		g.Group(children),
	)
}