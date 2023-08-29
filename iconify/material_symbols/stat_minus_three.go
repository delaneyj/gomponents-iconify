package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatMinusThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21.925l-6-6l1.4-1.4L12 19.1l4.6-4.575l1.4 1.4l-6 6Zm0-5.95l-6-6l1.4-1.4L12 13.15l4.6-4.575l1.4 1.4l-6 6Zm0-5.95l-6-6l1.4-1.4L12 7.2l4.6-4.575l1.4 1.4l-6 6Z"/>`),
		g.Group(children),
	)
}