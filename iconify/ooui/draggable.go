package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Draggable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 11h16v2H2zm0-4h16v2H2zm11 8H7l3 3zM7 5h6l-3-3z"/>`),
		g.Group(children),
	)
}