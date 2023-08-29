package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseMinimalisticLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5 9a7 7 0 0 1 14 0v6a7 7 0 1 1-14 0V9Z"/><path stroke-linecap="round" d="M12 5v3" opacity=".5"/></g>`),
		g.Group(children),
	)
}