package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipStartAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 15h2v-2H8v2m0-3h2V7H8v5M1 2v18h1v1h14v-1h1v-5h1v-1h1v-1h1v-1h1v-2h-1V9h-1V8h-1V7h-1V2h-1V1H2v1H1m2 1h12v5h1v1h1v1h1v2h-1v1h-1v1h-1v5H3V3Z"/>`),
		g.Group(children),
	)
}