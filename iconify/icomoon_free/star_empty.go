package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m16 6.204l-5.528-.803L8 .392L5.528 5.401L0 6.204l4 3.899l-.944 5.505L8 13.009l4.944 2.599L12 10.103l4-3.899zm-8 5.569l-3.492 1.836l.667-3.888L2.35 6.968l3.904-.567L8 2.864l1.746 3.537l3.904.567l-2.825 2.753l.667 3.888L8 11.773z"/>`),
		g.Group(children),
	)
}