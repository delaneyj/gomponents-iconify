package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-8H3l9-11l9 11h-5v8H8Zm2-2h4v-8h2.775L12 5.15L7.225 11H10v8Zm2-8Z"/>`),
		g.Group(children),
	)
}