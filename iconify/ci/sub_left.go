package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 5v8H7.83l2.58-2.59L9 9l-5 5l5 5l1.41-1.41L7.83 15H20V5h-2Z"/>`),
		g.Group(children),
	)
}