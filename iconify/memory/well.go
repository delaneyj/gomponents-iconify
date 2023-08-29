package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Well(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 21H3v-5H1v-2h20v2h-2v5M5 16v3h12v-3H5M2 7V5h1V4h1V3h1V2h1V1h10v1h1v1h1v1h1v1h1v2h-2v6h-2V7h-4v2h2v4H8V9h2V7H6v6H4V7H2m5-4v1H6v1h10V4h-1V3H7Z"/>`),
		g.Group(children),
	)
}