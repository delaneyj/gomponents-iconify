package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuBottomLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 4h2v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v2H7V4m2 4v5h5v-1h-1v-1h-1v-1h-1V9h-1V8H9Z"/>`),
		g.Group(children),
	)
}