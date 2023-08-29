package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 10v2H9v-2h1M6 2h10v1h1v15h2v2H3v-2h2V3h1V2m1 2v14h4V4H7m6 0v1h1V4h-1m1 1v1h1V5h-1m0 1h-1v1h1V6m0 1v1h1V7h-1m0 1h-1v1h1V8m0 1v1h1V9h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1m0 1h-1v1h1v-1m0 1v1h1v-1h-1Z"/>`),
		g.Group(children),
	)
}