package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpFilterAltOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.05 4H6.83l7.97 7.97zM2.81 2.81L1.39 4.22L10 13v7h4v-3.17l5.78 5.78l1.41-1.42z"/>`),
		g.Group(children),
	)
}