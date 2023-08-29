package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossRingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19.454 26.444C17.636 28.222 15.819 30 12.182 30C8.545 30 4 27.333 4 22s4.545-8 8.182-8c5.454 0 8.182 3.556 11.818 8c3.636 4.444 6.364 8 11.818 8C39.455 30 44 27.333 44 22s-4.545-8-8.182-8c-3.636 0-6.363 2.667-7.272 3.556"/>`),
		g.Group(children),
	)
}