package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronDownTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1024 1757l878-879l121 121l-999 999L25 999l121-121l878 879zm999-1611l-999 999L25 146L146 25l878 878l878-878l121 121z"/>`),
		g.Group(children),
	)
}