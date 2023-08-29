package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionRotaryFirstRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 18v6.586L15.871 14.457A5.924 5.924 0 0 0 17 11a6 6 0 1 0-7 5.91V28h2V16.91a5.957 5.957 0 0 0 2.455-1.04L24.585 26H18v2h10V18Zm-15-3a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}