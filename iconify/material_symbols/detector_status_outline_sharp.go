package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectorStatusOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.95 21L7.4 17.45l1.425-1.4l2.125 2.125l4.225-4.25L16.6 15.35L10.95 21ZM5 5v1h14V5H5Zm3.1 3l.3 1h7.2l.3-1H8.1Zm-1.15 3L6 8H3V3h18v5h-3l-1.15 3h-9.9ZM5 5v1v-1Z"/>`),
		g.Group(children),
	)
}