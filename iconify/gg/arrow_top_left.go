package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.475 5.495l.004 2l-4.557.01l9.603 9.585l-1.412 1.415l-9.63-9.61l.01 4.614l-2 .004l-.018-8l8-.018Z"/>`),
		g.Group(children),
	)
}