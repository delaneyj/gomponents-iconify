package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M215.39 124.94a8 8 0 0 1-1.73 8.72l-80 80a8 8 0 0 1-11.32 0l-80-80A8 8 0 0 1 48 120h60.69L42.34 53.66A8 8 0 0 1 48 40h160a8 8 0 0 1 5.66 13.66L147.31 120H208a8 8 0 0 1 7.39 4.94Z"/>`),
		g.Group(children),
	)
}