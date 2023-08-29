package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeBringForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 1664H0V0h1664v1664zM1536 128H128v1408h1408V128zm512 640v1280H768v-256h128v128h1024V896h-128V768h256z"/>`),
		g.Group(children),
	)
}