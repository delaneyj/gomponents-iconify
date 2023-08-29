package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartGroupedLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 12h2v9H2v-9Zm3 2h2v7H5v-7Zm11-6h2v13h-2V8Zm3 2h2v11h-2V10ZM9 2h2v19H9V2Zm3 2h2v17h-2V4Z"/>`),
		g.Group(children),
	)
}