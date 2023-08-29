package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Easy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 41c-5.48 3.003-10 3-16 1c-6.076-2.026-10-9-10-14c0-2.744 4.114-4.372 5.529-4.85a.687.687 0 0 0 .471-.648V15a3.5 3.5 0 1 1 7 0v-2.5a3.5 3.5 0 1 1 7 0V15a3.5 3.5 0 1 1 7 0V7.5a3.5 3.5 0 0 1 7 0v20.732c0 2.472-.592 4.928-2.038 6.934C37.492 37.205 35.357 39.708 33 41Z"/>`),
		g.Group(children),
	)
}