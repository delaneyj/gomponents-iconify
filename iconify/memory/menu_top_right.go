package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuTopRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M15 18h-2v-1h-1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V7h11v11m-2-4V9H8v1h1v1h1v1h1v1h1v1h1Z"/>`),
		g.Group(children),
	)
}