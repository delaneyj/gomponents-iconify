package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SouthEast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19v-2h6.6L4 5.4L5.4 4L17 15.6V9h2v10H9Z"/>`),
		g.Group(children),
	)
}