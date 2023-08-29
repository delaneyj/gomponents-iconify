package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolDoubleX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m10 4l14 14L38 4M10 44l14-14l14 14m5-34L29 24l14 14M4 10l14 14L4 38"/>`),
		g.Group(children),
	)
}