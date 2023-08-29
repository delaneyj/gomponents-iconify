package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStartArrowOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 19L2 12l11-7v6h9v2h-9v6Zm-2-3.65v-6.7L5.725 12L11 15.35ZM11 12Z"/>`),
		g.Group(children),
	)
}