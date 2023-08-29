package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Olympics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 9a3 3 0 1 0 6 0a3 3 0 1 0-6 0m12 0a3 3 0 1 0 6 0a3 3 0 1 0-6 0M9 9a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/><path d="M6 15a3 3 0 1 0 6 0a3 3 0 1 0-6 0m6 0a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/></g>`),
		g.Group(children),
	)
}