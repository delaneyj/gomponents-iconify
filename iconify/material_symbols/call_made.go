package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallMade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.4 20L4 18.6L15.6 7H9V5h10v10h-2V8.4L5.4 20Z"/>`),
		g.Group(children),
	)
}