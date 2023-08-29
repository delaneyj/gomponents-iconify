package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.63 5.643a9 9 0 0 0 12.742 12.715m1.674-2.29A9.03 9.03 0 0 0 20.8 14a1 1 0 0 0-1-1H17m-4 0a2 2 0 0 1-2-2m0-4V4a.9.9 0 0 0-1-.8a9 9 0 0 0-2.057.749M15 3.5A9 9 0 0 1 20.5 9H16a1 1 0 0 1-1-1V3.5M3 3l18 18"/>`),
		g.Group(children),
	)
}