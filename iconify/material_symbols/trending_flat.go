package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingFlat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.5 16.5l-1.425-1.4l2.1-2.1H3v-2h15.175L16.1 8.9l1.425-1.4L22 12l-4.5 4.5Z"/>`),
		g.Group(children),
	)
}