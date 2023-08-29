package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncreaseDecreaseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.005 3.003h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1v-16a1 1 0 0 1 1-1Zm1 2v14h16v-14h-16Zm5 6h2v2h-2v2h-2v-2h-2v-2h2v-2h2v2Zm4 0h6v2h-6v-2Z"/>`),
		g.Group(children),
	)
}