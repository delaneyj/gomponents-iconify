package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Level(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 42L4 18.5L9.695 6h28.61L44 18.5L24 42Z"/><path d="m32 18l-8 9l-8-9"/></g>`),
		g.Group(children),
	)
}