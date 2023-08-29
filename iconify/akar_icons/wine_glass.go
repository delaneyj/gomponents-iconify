package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13c6 0 8-4.477 8-10H4c0 5.523 2 10 8 10Zm0 0v7M5 8h14M8 22h8l-4-2l-4 2Z"/>`),
		g.Group(children),
	)
}