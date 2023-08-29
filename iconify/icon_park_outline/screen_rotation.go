package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 24c0 11.046 8.954 20 20 20l-5-5m25-15c0-11.046-8.954-20-20-20l5 5m1 32L7 18L18 7l23 23l-11 11Z"/>`),
		g.Group(children),
	)
}