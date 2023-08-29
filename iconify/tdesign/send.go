package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Send(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.292 1.665L24.002 12L.293 22.336L3.94 12L.292 1.665ZM5.708 13l-2 5.665L18.999 12L3.708 5.336l2 5.664H11v2H5.708Z"/>`),
		g.Group(children),
	)
}