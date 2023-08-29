package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SweepOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 18v-2h6v2h-6Zm-3.95 0L.375 12.325L1.8 10.9l4.25 4.25L15.2 6l1.425 1.425L6.05 18ZM14 14v-2h6v2h-6Zm4-4V8h6v2h-6Z"/>`),
		g.Group(children),
	)
}