package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleChevronStartLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m9 2l1.3 1.3L3.7 10l6.6 6.7L9 18l-8-8l8-8zm8.5 0L19 3.3L12.2 10l6.7 6.7l-1.4 1.3l-8-8l8-8z"/>`),
		g.Group(children),
	)
}