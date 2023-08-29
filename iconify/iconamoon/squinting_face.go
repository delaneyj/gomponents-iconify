package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquintingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path fill="currentColor" d="M12 16c1.48 0 2.773-.804 3.465-2h-6.93A3.998 3.998 0 0 0 12 16Z"/><path d="m16 10.5l-2-1l2-1m-8 2l2-1l-2-1"/></g>`),
		g.Group(children),
	)
}