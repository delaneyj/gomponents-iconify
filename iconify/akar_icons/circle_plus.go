package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M12 8v4m0 0v4m0-4h4m-4 0H8"/><circle cx="12" cy="12" r="10"/></g>`),
		g.Group(children),
	)
}