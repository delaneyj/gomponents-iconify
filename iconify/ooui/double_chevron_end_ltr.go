package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronEndLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 2L9.7 3.3l6.6 6.7l-6.6 6.7L11 18l8-8zM2.5 2L1 3.3L7.8 10l-6.7 6.7L2.5 18l8-8z"/>`),
		g.Group(children),
	)
}