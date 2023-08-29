package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatMinusTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 19l-6-6l1.4-1.4l4.6 4.575l4.6-4.575L18 13l-6 6Zm0-6L6 7l1.4-1.4l4.6 4.575L16.6 5.6L18 7l-6 6Z"/>`),
		g.Group(children),
	)
}