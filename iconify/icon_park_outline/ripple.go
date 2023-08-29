package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ripple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 36s5-3 10-3c7.298 0 12.702 6 20 6c5 0 10-3 10-3M4 24s5-3 10-3c7.298 0 12.702 6 20 6c5 0 10-3 10-3M4 12s5-3 10-3c7.298 0 12.702 6 20 6c5 0 10-3 10-3"/>`),
		g.Group(children),
	)
}