package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronRightTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m999 25l999 999l-999 999l-121-121l879-878l-879-878L999 25zm-853 0l999 999l-999 999l-121-121l878-878L25 146L146 25z"/>`),
		g.Group(children),
	)
}