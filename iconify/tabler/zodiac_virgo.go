package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZodiacVirgo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 4a2 2 0 0 1 2 2v9m0-9a2 2 0 0 1 4 0v9m0-9a2 2 0 0 1 4 0v10a7 5 0 0 0 7 5"/><path d="M12 21a7 5 0 0 0 7-5v-2a3 3 0 0 0-6 0"/></g>`),
		g.Group(children),
	)
}