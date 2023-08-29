package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartColum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v2h15v6H4v2h11v6H4v2h18v2H2V2h2Zm0 14h9v-2H4v2Zm0-8h13V6H4v2Z"/>`),
		g.Group(children),
	)
}