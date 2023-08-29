package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsScrollV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.414 10.11l.004.003l-1.414 1.414l-.004-.004l-.004.004l-1.414-1.414l.004-.004L7.76 7.284L9.175 5.87L12 8.695l2.825-2.825l1.414 1.414l-2.825 2.825ZM12 15.305l2.825 2.825l1.414-1.414l-2.825-2.825l.004-.004l-1.414-1.414l-.004.004l-.004-.004l-1.414 1.414l.004.004l-2.825 2.825l1.414 1.414L12 15.305Z"/>`),
		g.Group(children),
	)
}