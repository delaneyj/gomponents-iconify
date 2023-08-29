package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateNot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4v16l14.2-7c.42 1.19 1.54 2 2.8 2a3 3 0 0 0 3-3a3 3 0 0 0-3-3c-1.26 0-2.38.81-2.8 2L2 4m2 3.3l9.7 4.7L4 16.7V7.3M19 11c.5 0 1 .5 1 1s-.5 1-1 1a1 1 0 0 1-1-1c0-.5.5-1 1-1Z"/>`),
		g.Group(children),
	)
}