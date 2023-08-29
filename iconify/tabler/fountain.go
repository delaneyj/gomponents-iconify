package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fountain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 16v-5a2 2 0 1 0-4 0m10 5v-5a2 2 0 1 1 4 0"/><path d="M12 16V6a3 3 0 0 1 6 0M6 6a3 3 0 0 1 6 0M3 16h18v2a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-2z"/></g>`),
		g.Group(children),
	)
}