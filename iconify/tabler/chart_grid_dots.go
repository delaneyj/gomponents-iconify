package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartGridDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 6a2 2 0 1 0 4 0a2 2 0 1 0-4 0M4 12a2 2 0 1 0 4 0a2 2 0 1 0-4 0m0 6a2 2 0 1 0 4 0a2 2 0 1 0-4 0m12 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-8 0h8m2 2v1m0-18v1M6 20v1m0-11V3m6 0v18m6-13v8M8 12h13m0-6h-1m-4 0H3m0 6h1m16 6h1M3 18h1m2-4v2"/>`),
		g.Group(children),
	)
}