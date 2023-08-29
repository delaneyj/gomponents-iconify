package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightUpRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M9 7h4v1h1v1h1v4h-1v1h-1v1h-1v7h-2v-7H9v-1H8v-1H7v-1H0v-2h7V9h1V8h1V7m4 3h-1V9h-2v1H9v2h1v1h2v-1h1v-2Z"/>`),
		g.Group(children),
	)
}