package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartSankey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 5H4V2H2v20h20v-2H4V9c4.09 0 6.13 2 8.29 4.21S17.09 18 22 18v-2c-4.09 0-6.13-2-8.29-4.21S8.91 7 4 7h18Z"/>`),
		g.Group(children),
	)
}