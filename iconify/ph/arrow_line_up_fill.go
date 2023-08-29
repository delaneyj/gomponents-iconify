package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M205.66 138.34A8 8 0 0 1 200 152h-64v72a8 8 0 0 1-16 0v-72H56a8 8 0 0 1-5.66-13.66l72-72a8 8 0 0 1 11.32 0ZM216 32H40a8 8 0 0 0 0 16h176a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}