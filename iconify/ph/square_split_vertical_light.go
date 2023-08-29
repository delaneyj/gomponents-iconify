package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareSplitVerticalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 42H56a14 14 0 0 0-14 14v144a14 14 0 0 0 14 14h144a14 14 0 0 0 14-14V56a14 14 0 0 0-14-14ZM56 54h144a2 2 0 0 1 2 2v66H54V56a2 2 0 0 1 2-2Zm144 148H56a2 2 0 0 1-2-2v-66h148v66a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}