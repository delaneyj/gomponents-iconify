package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NorthWest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.6 20L7 8.4V15H5V5h10v2H8.4L20 18.6L18.6 20Z"/>`),
		g.Group(children),
	)
}