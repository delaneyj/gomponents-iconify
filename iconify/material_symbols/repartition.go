package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repartition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 14v-2h14q.825 0 1.413-.588T19 10q0-.825-.588-1.413T17 8H5.825l1.6 1.575L6 11L2 7l4-4l1.425 1.4l-1.6 1.6H17q1.65 0 2.825 1.175T21 10q0 1.65-1.175 2.825T17 14H3Zm0 8v-6h18v6H3Zm2-2h3.325v-2H5v2Zm5.325 0h3.325v-2h-3.325v2Zm5.35 0H19v-2h-3.325v2Z"/>`),
		g.Group(children),
	)
}