package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.5 8H15s0 0 0 0s5 0 5 4.706C20 18 15 18 15 18H6.286"/><path d="M7.5 11.5L4 8l3.5-3.5"/></g>`),
		g.Group(children),
	)
}