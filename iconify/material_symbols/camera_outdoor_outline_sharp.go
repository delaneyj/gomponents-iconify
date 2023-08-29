package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOutdoorOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18v-6h6v2l2-1.05v4.1L18 16v2h-6Zm-8 3V9l8-6l8 6v2h-2v-1l-6-4.5L6 10v9h14v2H4Zm8-8.75Z"/>`),
		g.Group(children),
	)
}