package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBoxPlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 5v2h-3v3h-2V7h-3V5h3V2h2v3h3m-3 14H5V5h6V3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2v-6h-2v6m-4-6v4h2v-4h-2m-4 4h2V9h-2v8m-2 0v-6H7v6h2Z"/>`),
		g.Group(children),
	)
}