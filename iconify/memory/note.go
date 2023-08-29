package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Note(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 3v1h1v1h1v1h1v1h1v1h1v1h1v9h-1v1H2v-1H1V4h1V3h13m0 3h-1v4h4V9h-1V8h-1V7h-1V6M3 5v12h16v-5h-6v-1h-1V5H3Z"/>`),
		g.Group(children),
	)
}