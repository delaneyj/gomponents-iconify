package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartInfographic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7a4 4 0 1 0 8 0a4 4 0 1 0-8 0"/><path d="M7 3v4h4M9 17v4m8-7v7m-4-8v8m8-9v9"/></g>`),
		g.Group(children),
	)
}