package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartCombo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h2v-6h6v6h2v-9h6v9h2v2H2V2h2Zm14 18v-7h-2v7h2Zm-8 0v-4H8v4h2Zm6.673-16.273L21.246 8.3l-1.414 1.414l-3.163-3.163l-6.782 6.74l-4.594-4.594l1.414-1.415l3.184 3.185l6.782-6.74Z"/>`),
		g.Group(children),
	)
}