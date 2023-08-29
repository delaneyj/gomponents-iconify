package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteMediumSmallSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#d0d2d3" d="M49.717 44.915a3.551 3.551 0 0 1-3.549 3.552H20.332a3.55 3.55 0 0 1-3.548-3.552V19.083a3.55 3.55 0 0 1 3.548-3.551h25.836a3.55 3.55 0 0 1 3.549 3.551v25.832"/>`),
		g.Group(children),
	)
}