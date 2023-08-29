package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordMinimalisticLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M22 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-12 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/><path stroke-linecap="round" d="M6 16h12"/></g>`),
		g.Group(children),
	)
}