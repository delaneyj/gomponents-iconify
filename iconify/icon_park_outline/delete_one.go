package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="m15 12l1.2-7h15.6l1.2 7"/><path stroke-linecap="round" d="M6 12h36"/><path stroke-linecap="round" stroke-linejoin="round" d="m37 12l-2 31H13l-2-31h26Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M19 35h10"/></g>`),
		g.Group(children),
	)
}