package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 341v683h-171V633l-768 768l-341-342l-665 665l-121-120l786-786l341 341l648-647h-392V341h683z"/>`),
		g.Group(children),
	)
}