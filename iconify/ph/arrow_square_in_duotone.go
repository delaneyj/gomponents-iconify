package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareInDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 48v128a8 8 0 0 1-8 8H80a8 8 0 0 1-8-8V48a8 8 0 0 1 8-8h128a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M128 136v64a8 8 0 0 1-16 0v-44.68l-66.34 66.34a8 8 0 0 1-11.32-11.32L100.68 144H56a8 8 0 0 1 0-16h64a8 8 0 0 1 8 8Zm80-104H80a16 16 0 0 0-16 16v48a8 8 0 0 0 16 0V48h128v128h-48a8 8 0 0 0 0 16h48a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16Z"/></g>`),
		g.Group(children),
	)
}