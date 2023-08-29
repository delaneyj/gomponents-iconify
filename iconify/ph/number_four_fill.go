package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFourFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 24H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16Zm-32 128h-8v32a8 8 0 0 1-16 0v-32H88a8 8 0 0 1-7.43-11l32-80a8 8 0 0 1 14.86 6l-27.61 69H144v-32a8 8 0 0 1 16 0v32h8a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}