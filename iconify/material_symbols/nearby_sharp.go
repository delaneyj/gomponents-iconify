package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearbySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.4L7.6 12L12 7.6l4.4 4.4l-4.4 4.4Zm0 6.375L1.225 12L12 1.225L22.775 12L12 22.775Zm0-3.575l7.2-7.2L12 4.8L4.8 12l7.2 7.2Z"/>`),
		g.Group(children),
	)
}