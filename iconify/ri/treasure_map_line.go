package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreasureMapLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.935 7.204l-6-3L4 6.319v12.648l5.065-2.17l6 3L20 17.68V5.033l-5.065 2.17ZM2 5l7-3l6 3l6.303-2.701a.5.5 0 0 1 .697.46V19l-7 3l-6-3l-6.303 2.701a.5.5 0 0 1-.697-.46V5Zm4 6h2v2H6v-2Zm4 0h2v2h-2v-2Zm5.998-.063L17.235 9.7l1.061 1.06l-1.237 1.238l1.237 1.238l-1.06 1.06l-1.238-1.237l-1.237 1.237l-1.061-1.06l1.237-1.238l-1.237-1.237L14.76 9.7l1.238 1.237Z"/>`),
		g.Group(children),
	)
}