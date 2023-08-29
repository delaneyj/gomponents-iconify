package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronUpTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1024 291l-878 879l-121-121l999-999l999 999l-121 121l-878-879zM25 1902l999-999l999 999l-121 121l-878-878l-878 878l-121-121z"/>`),
		g.Group(children),
	)
}