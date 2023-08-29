package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSignalGraphCircleCircleStatsGraphLineBeatHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7.01h4L6 4.5L7.5 10l2-2.99h4M12.77 4A6.51 6.51 0 0 0 1.23 4m0 6a6.51 6.51 0 0 0 11.54 0"/>`),
		g.Group(children),
	)
}