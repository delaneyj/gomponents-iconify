package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Goblet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 44H13m10-16v16m12-28c0 6.5-5.373 12-12 12s-12-5.373-12-12c0-6.5 4-12 4-12h16s4 5.5 4 12Z"/>`),
		g.Group(children),
	)
}