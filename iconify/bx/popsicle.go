package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Popsicle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 2a7 7 0 0 0-4.94 2l-7.78 7.82a1 1 0 0 0 0 1.41l3.54 3.54l-3.54 3.53l1.42 1.42l3.53-3.54l3.54 3.54a1 1 0 0 0 1.41 0L20 13.94A7 7 0 0 0 15 2zm3.54 10.54l-7.07 7.06l-2.82-2.83l-1.42-1.42l-2.83-2.83l7.07-7.07a5 5 0 0 1 7.08 7.08z"/>`),
		g.Group(children),
	)
}