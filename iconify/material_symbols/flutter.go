package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flutter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.9 15.375L2.5 12l11-11h6.775L5.9 15.375ZM13.5 23l-5.925-5.925L13.5 11.15h6.775l-5.925 5.925L20.275 23H13.5Z"/>`),
		g.Group(children),
	)
}