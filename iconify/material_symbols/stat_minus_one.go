package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatMinusOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 15.375l-6-6l1.4-1.4L12 12.55l4.6-4.575l1.4 1.4l-6 6Z"/>`),
		g.Group(children),
	)
}