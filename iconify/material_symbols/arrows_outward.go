package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 17l-1.4-1.4l2.575-2.6H13v-2h5.175L15.6 8.4L17 7l5 5l-5 5ZM7 17l-5-5l5-5l1.4 1.4L5.825 11H11v2H5.825L8.4 15.6L7 17Z"/>`),
		g.Group(children),
	)
}