package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddFavoriteFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1860 832l-519 399l90 286l-23 19h-128v78l-39 28l-217-167l-519 399l202-643l-519-399h643l193-618l193 618h643z"/>`),
		g.Group(children),
	)
}