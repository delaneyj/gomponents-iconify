package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GestureDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.57 19.5a3 3 0 1 0 6 0V16h1.46a3 3 0 0 0 1.511-.409l4.13-2.409a3 3 0 0 0 1.334-3.54l-2.03-6.09A3 3 0 0 0 17.129 1.5H9.055a3 3 0 0 0-2.378 1.17L1.841 8.959l1.017 1.527a2 2 0 0 0 2.098.843l2.614-.581V19.5Zm3 1a1 1 0 0 1-1-1V8.253L4.522 9.375L4.3 9.042L8.263 3.89a1 1 0 0 1 .792-.39h8.074a1 1 0 0 1 .949.684l2.03 6.09a1 1 0 0 1-.445 1.18l-4.13 2.41a1 1 0 0 1-.503.136h-3.46v5.5a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}