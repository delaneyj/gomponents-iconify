package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TogoCup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 5.4A2.4 2.4 0 0 1 5.4 3h13.2A2.4 2.4 0 0 1 21 5.4v0a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6v0ZM5 6h14l-1.555 14.218A2 2 0 0 1 15.457 22H8.543a2 2 0 0 1-1.988-1.782L5 6Z"/><path d="m6.313 18l-.875-8h13.125l-.875 8H6.313Z"/></g>`),
		g.Group(children),
	)
}