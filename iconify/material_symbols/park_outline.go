package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.95 22h-3.9v-4H3l4-6H5l7-10l7 10h-2l4 6h-7.05v4Zm-7.2-6h4h-1.9h6.3h-1.9h4h-10.5Zm0 0h10.5l-4-6h1.9L12 5.5L8.85 10h1.9l-4 6Z"/>`),
		g.Group(children),
	)
}