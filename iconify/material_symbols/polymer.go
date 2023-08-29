package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polymer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.95 20L.5 12L5 4h4l-4.5 8l2.6 4.65L14.9 4H19l4.5 8l-4.5 8h-4l4.5-8l-2.6-4.6L9.15 20h-4.2Z"/>`),
		g.Group(children),
	)
}