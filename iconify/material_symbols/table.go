package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3ZM5 9h14V5H5v4Zm5.325 5h3.35v-3h-3.35v3Zm0 5h3.35v-3h-3.35v3ZM5 14h3.325v-3H5v3Zm10.675 0H19v-3h-3.325v3ZM5 19h3.325v-3H5v3Zm10.675 0H19v-3h-3.325v3Z"/>`),
		g.Group(children),
	)
}