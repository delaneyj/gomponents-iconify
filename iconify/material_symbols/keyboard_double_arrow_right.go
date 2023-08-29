package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardDoubleArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.575 12L5 7.4L6.4 6l6 6l-6 6L5 16.6L9.575 12Zm6.6 0L11.6 7.4L13 6l6 6l-6 6l-1.4-1.4l4.575-4.6Z"/>`),
		g.Group(children),
	)
}