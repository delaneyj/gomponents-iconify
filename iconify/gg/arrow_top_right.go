package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.525 5.495l-.004 2l4.557.01l-9.603 9.585l1.413 1.415l9.63-9.61l-.012 4.614l2 .004l.02-8l-8-.018Z"/>`),
		g.Group(children),
	)
}