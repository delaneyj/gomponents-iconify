package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeTwoLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m16 11l1.697 1.527c1.542 1.388 2.313 2.082 2.313 2.973c0 .89-.771 1.585-2.314 2.973L16 20M13.987 5l-1.994 7.5L10 20M8 4.83L6.304 6.356C4.76 7.745 3.99 8.44 3.99 9.33c0 .89.771 1.585 2.314 2.973L8 13.83"/>`),
		g.Group(children),
	)
}