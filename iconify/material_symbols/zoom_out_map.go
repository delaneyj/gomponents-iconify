package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-6h2v2.6l3.1-3.1l1.4 1.4L6.4 19H9v2H3Zm12 0v-2h2.6l-3.1-3.1l1.4-1.4l3.1 3.1V15h2v6h-6ZM8.1 9.5L5 6.4V9H3V3h6v2H6.4l3.1 3.1l-1.4 1.4Zm7.8 0l-1.4-1.4L17.6 5H15V3h6v6h-2V6.4l-3.1 3.1Z"/>`),
		g.Group(children),
	)
}