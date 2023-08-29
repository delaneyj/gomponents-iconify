package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanZoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-6h2v2.6l3.1-3.1l1.4 1.4L6.4 19H9v2H3ZM15.9 9.5l-1.4-1.4L17.6 5H15V3h6v6h-2V6.4l-3.1 3.1Z"/>`),
		g.Group(children),
	)
}