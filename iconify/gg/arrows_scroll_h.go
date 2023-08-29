package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsScrollH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.305 12l2.825-2.825l-1.414-1.414l-2.825 2.825l-.004-.004l-1.414 1.414l.004.004l-.004.004l1.414 1.414l.004-.004l2.825 2.825l1.414-1.414L15.305 12Zm-5.195-1.414l.003-.004l1.414 1.414l-.004.004l.004.004l-1.414 1.414l-.004-.004l-2.825 2.825l-1.414-1.414L8.695 12L5.87 9.175l1.414-1.414l2.825 2.825Z"/>`),
		g.Group(children),
	)
}