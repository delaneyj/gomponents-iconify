package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartGantt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5h8V2h2v20h-2v-4H6v-3h4v-2H4v-3h6V8H2V5m12 0h3v3h-3V5m0 5h5v3h-5v-3m0 5h8v3h-8v-3Z"/>`),
		g.Group(children),
	)
}