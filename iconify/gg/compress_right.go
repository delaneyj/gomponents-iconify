package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompressRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.567 8.03l6.343-6.26l1.405 1.423l-6.323 6.24l3.57.015l-.009 2l-7-.028l.028-7l2 .008l-.014 3.601Zm-6.588 6.513l-3.57.003l-.002-2l7-.006l.006 7l-2 .002l-.003-3.602l-6.314 6.29l-1.411-1.416l6.294-6.271Z"/>`),
		g.Group(children),
	)
}