package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineEndArrowNotch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 19l3.425-6H2v-2h12.425L11 5l11 7l-11 7Z"/>`),
		g.Group(children),
	)
}