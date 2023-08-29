package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardReader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5 19V3h14v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2Z"/><path d="M5 6H3.5a1.5 1.5 0 1 1 0-3h17a1.5 1.5 0 0 1 0 3H19"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 3v18"/></g>`),
		g.Group(children),
	)
}