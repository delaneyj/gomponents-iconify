package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnalyzeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 11a8.1 8.1 0 0 0-6.986-6.918a8.086 8.086 0 0 0-4.31.62M6.321 6.31A8.089 8.089 0 0 0 4.995 8M4 13a8.1 8.1 0 0 0 13.687 4.676M20 16a1 1 0 0 0-1-1"/><path d="M4 8a1 1 0 1 0 2 0a1 1 0 1 0-2 0m5.888 1.87a3 3 0 1 0 4.233 4.252m.595-3.397A3.012 3.012 0 0 0 13.29 9.29M3 3l18 18"/></g>`),
		g.Group(children),
	)
}