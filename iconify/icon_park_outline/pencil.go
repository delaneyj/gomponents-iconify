package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m31 8.999l8 8m-31 15L36 4l8 7.999l-28 28l-10 2l2-10Zm23-23l8 8m-30 15l7 7m-3-4l22-22"/>`),
		g.Group(children),
	)
}