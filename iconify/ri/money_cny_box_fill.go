package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCnyBoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.005 3.003h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1v-16a1 1 0 0 1 1-1Zm10 10v-1h3v-2h-2.586l2.121-2.121l-1.414-1.415l-2.121 2.122l-2.121-2.122l-1.415 1.415l2.122 2.12H8.005v2h3v1h-3v2h3v2h2v-2h3v-2h-3Z"/>`),
		g.Group(children),
	)
}