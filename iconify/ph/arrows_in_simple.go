package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.66 53.66L163.31 104H192a8 8 0 0 1 0 16h-48a8 8 0 0 1-8-8V64a8 8 0 0 1 16 0v28.69l50.34-50.35a8 8 0 0 1 11.32 11.32ZM112 136H64a8 8 0 0 0 0 16h28.69l-50.35 50.34a8 8 0 0 0 11.32 11.32L104 163.31V192a8 8 0 0 0 16 0v-48a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}