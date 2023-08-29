package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 13v1h12v-1h-1v-1h-1v-1h-1v-1h-1V9h-1V8h-2v1H9v1H8v1H7v1H6v1"/>`),
		g.Group(children),
	)
}