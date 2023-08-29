package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m16 6.204l-5.528-.803L8 .392L5.528 5.401L0 6.204l4 3.899l-.944 5.505L8 13.009l4.944 2.599L12 10.103l4-3.899z"/>`),
		g.Group(children),
	)
}