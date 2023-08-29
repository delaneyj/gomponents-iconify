package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Water(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 14c2.077 0 4.5-1.288 4.5-4.568C12 7.207 8.538 2.288 7.5 1C6.577 2.288 3 7.09 3 9.432C3 12.712 5.423 14 7.5 14Z"/>`),
		g.Group(children),
	)
}