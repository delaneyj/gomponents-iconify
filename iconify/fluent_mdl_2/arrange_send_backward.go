package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeSendBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1536h512v128H0V0h1664v640h-128V128H128v1408zm1920-768v1280H768V768h1280zm-128 128H896v1024h1024V896z"/>`),
		g.Group(children),
	)
}