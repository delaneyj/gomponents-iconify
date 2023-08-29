package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GobletFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19v-5.111L3 5V3h18v2l-8 8.889V19h5v2H6v-2h5ZM7.49 7h9.02l1.8-2H5.69l1.8 2Z"/>`),
		g.Group(children),
	)
}