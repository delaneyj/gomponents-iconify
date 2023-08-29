package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingUpThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m18 5l3 3l-3 3"/><path d="M3 18h2.397a5 5 0 0 0 4.096-2.133l4.014-5.734A5 5 0 0 1 17.603 8H21"/></g>`),
		g.Group(children),
	)
}