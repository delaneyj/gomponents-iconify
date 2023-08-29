package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMonth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h2V0h2v2h8V0h2v2h2v1h1v16h-1v1M4 4v2h14V4H4m0 4v10h14V8H4m10 6h2v2h-2v-2m-4 0h2v2h-2v-2m-4 0h2v2H6v-2m0-4h2v2H6v-2m4 0h2v2h-2v-2m4 0h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}