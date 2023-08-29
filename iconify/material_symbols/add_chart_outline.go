package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddChartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h9v2H5v14h14v-9h2v9q0 .825-.588 1.413T19 21H5Zm2-4h2v-7H7v7Zm4 0h2V7h-2v10Zm4 0h2v-4h-2v4Zm2-8V7h-2V5h2V3h2v2h2v2h-2v2h-2Zm-5 3Z"/>`),
		g.Group(children),
	)
}