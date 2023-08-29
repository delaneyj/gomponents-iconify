package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDoubleLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.299 7.76l-5.019 4.88l-1.394-1.434l2.436-2.368l-6.02.015a2.4 2.4 0 0 0-2.394 2.406l.014 5.9l2.268-2.256l1.41 1.418l-4.962 4.937l-4.937-4.962l1.418-1.41L6.522 17.3l-.014-6.036a4.8 4.8 0 0 1 4.788-4.812l5.928-.014l-2.239-2.303l1.434-1.394L21.3 7.76Z"/>`),
		g.Group(children),
	)
}