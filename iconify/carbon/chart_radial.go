package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartRadial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 30A14.016 14.016 0 0 1 2 16h2A12 12 0 1 0 16 4V2a14 14 0 0 1 0 28Z"/><path fill="currentColor" d="M16 26A10.011 10.011 0 0 1 6 16h2a8 8 0 1 0 8-8V6a10 10 0 0 1 0 20Z"/><path fill="currentColor" d="M16 22a6.007 6.007 0 0 1-6-6h2a4 4 0 1 0 4-4v-2a6 6 0 0 1 0 12Z"/>`),
		g.Group(children),
	)
}