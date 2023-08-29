package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartWaterfall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h2v18h18v2H2V2m15 0h3v16h-3V2M6 11h3v7H6v-7m7-8h3v4h-3V3m-3 5h3v4h-3V8Z"/>`),
		g.Group(children),
	)
}