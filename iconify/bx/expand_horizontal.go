package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.29 5.64L1.93 12l6.36 6.36l1.42-1.41L4.76 12l4.95-4.95l-1.42-1.41zm6 1.41L19.24 12l-4.95 4.95l1.42 1.41L22.07 12l-6.36-6.36l-1.42 1.41z"/>`),
		g.Group(children),
	)
}