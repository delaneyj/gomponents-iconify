package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipStartText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 8V6h8v2H5m0 8v-2h6v2H5m0-4v-2h8v2H5m-4 8V2h1V1h14v1h1v5h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v5h-1v1H2v-1H1m2-1h12v-5h1v-1h1v-1h1v-2h-1V9h-1V8h-1V3H3v16Z"/>`),
		g.Group(children),
	)
}