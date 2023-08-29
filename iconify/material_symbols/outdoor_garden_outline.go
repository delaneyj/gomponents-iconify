package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutdoorGardenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V6l4-3l3 2.25L12 3l3 2.25L18 3l4 3v15H2Zm2-2h4V7L6 5.5L4 7v12Zm6 0h4V7l-2-1.5L10 7v12Zm6 0h4V7l-2-1.5L16 7v12Z"/>`),
		g.Group(children),
	)
}