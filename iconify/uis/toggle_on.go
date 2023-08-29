package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 6.5h-9C4.5 6.5 2 9 2 12s2.5 5.5 5.5 5.5h9c3 0 5.5-2.5 5.5-5.5s-2.5-5.5-5.5-5.5zm0 8c-1.4 0-2.5-1.1-2.5-2.5s1.1-2.5 2.5-2.5S19 10.6 19 12s-1.1 2.5-2.5 2.5z"/>`),
		g.Group(children),
	)
}