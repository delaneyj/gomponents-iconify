package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightDownRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 9v4h-1v1h-1v1H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1V7h1V0h2v7h1v1h1v1h1m-3 4v-1h1v-2h-1V9h-2v1H9v2h1v1h2Z"/>`),
		g.Group(children),
	)
}