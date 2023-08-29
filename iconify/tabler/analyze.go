package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analyze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 11a8.1 8.1 0 0 0-6.986-6.918A8.095 8.095 0 0 0 4.995 8M4 13a8.1 8.1 0 0 0 15 3"/><path d="M18 16a1 1 0 1 0 2 0a1 1 0 1 0-2 0M4 8a1 1 0 1 0 2 0a1 1 0 1 0-2 0m5 4a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/></g>`),
		g.Group(children),
	)
}