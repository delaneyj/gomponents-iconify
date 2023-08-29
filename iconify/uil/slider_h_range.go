package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SliderHRange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 11h-1.184a2.982 2.982 0 0 0-5.632 0H9.816a2.982 2.982 0 0 0-5.632 0H3a1 1 0 0 0 0 2h1.184a2.982 2.982 0 0 0 5.632 0h4.368a2.982 2.982 0 0 0 5.632 0H21a1 1 0 0 0 0-2ZM7 13a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1Zm10 0a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}