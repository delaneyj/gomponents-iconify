package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsAltSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22.8L1.2 12L12 1.2L22.8 12L12 22.8Zm0-5.8l5-5l-5-5l-1.4 1.4l2.55 2.6H7v2h6.15l-2.55 2.6L12 17Z"/>`),
		g.Group(children),
	)
}