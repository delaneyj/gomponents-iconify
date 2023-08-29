package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19.5 12.572l-3 2.928M11 19a8916.99 8916.99 0 0 0-6.5-6.428A5 5 0 1 1 12 6.006a5 5 0 1 1 7.5 6.572M15 19l2 2l4-4"/>`),
		g.Group(children),
	)
}