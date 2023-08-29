package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Label(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 4h13v1h1v1h1v1h1v1h1v1h1v1h1v2h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1H2v-1H1V5h1V4m14 9h1v-1h1v-2h-1V9h-1V8h-1V7h-1V6H3v10h11v-1h1v-1h1v-1Z"/>`),
		g.Group(children),
	)
}