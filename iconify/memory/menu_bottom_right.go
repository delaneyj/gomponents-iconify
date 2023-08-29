package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 15v-2h1v-1h1v-1h1v-1h1V9h1V8h1V7h1V6h1V5h1V4h2v11H4m4-2h5V8h-1v1h-1v1h-1v1H9v1H8v1Z"/>`),
		g.Group(children),
	)
}