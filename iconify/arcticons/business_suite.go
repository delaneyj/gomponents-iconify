package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessSuite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.33 34.669A21.51 21.51 0 0 1 8.796 8.797L24 24L8.797 39.203A21.497 21.497 0 1 0 13.33 5.33"/>`),
		g.Group(children),
	)
}