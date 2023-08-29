package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuadrangularPyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24L24 4L4 24l20 20l20-20ZM4 24h40"/><path d="m24 44l-6-20l6-20m0 40l6-20l-6-20"/></g>`),
		g.Group(children),
	)
}