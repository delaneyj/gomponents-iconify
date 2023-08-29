package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M8 18H6a2 2 0 0 1-2-2V7h16v9a2 2 0 0 1-2 2h-2M8 3h8v4H8V3Z"/><path stroke-linecap="round" d="M12 11H8"/><path d="M8 15h8v6H8v-6Z"/></g>`),
		g.Group(children),
	)
}