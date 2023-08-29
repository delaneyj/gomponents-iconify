package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M50.34 117.66A8 8 0 0 1 56 104h64V32a8 8 0 0 1 16 0v72h64a8 8 0 0 1 5.66 13.66l-72 72a8 8 0 0 1-11.32 0ZM216 208H40a8 8 0 0 0 0 16h176a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}