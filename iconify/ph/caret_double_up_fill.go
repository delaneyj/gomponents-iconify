package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDoubleUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.66 202.34A8 8 0 0 1 208 216H48a8 8 0 0 1-5.66-13.66L108.69 136H48a8 8 0 0 1-5.66-13.66l80-80a8 8 0 0 1 11.32 0l80 80A8 8 0 0 1 208 136h-60.69Z"/>`),
		g.Group(children),
	)
}