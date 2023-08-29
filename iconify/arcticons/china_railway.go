package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChinaRailway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.092 34.552a15.53 15.53 0 1 1 17.816 0M20.551 4.677h7.128"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.614 18.563h12.078v5.346a19.272 19.272 0 0 0-6.04 2.772c-1.125-.842-2.106-1.848-6.038-2.772Zm6.04 8.118v10.692c1.544 3.09 7.907 4.011 12.375 4.95H10.981c4.459-.794 10.357-1.429 12.672-4.95"/>`),
		g.Group(children),
	)
}