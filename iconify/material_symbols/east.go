package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func East(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 19l-1.425-1.4l4.6-4.6H2v-2h16.175L13.6 6.4L15 5l7 7l-7 7Z"/>`),
		g.Group(children),
	)
}