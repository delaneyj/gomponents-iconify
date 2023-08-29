package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Device(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 1h18v1h.94v18H20v1H2v-1h-.94V2H2V1m1 2v16h16V3H3m1 1h14v8H4V4m1 10h3v3H5v-3m7 1h2v2h-2v-2m3-1h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}