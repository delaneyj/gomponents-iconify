package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Post(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13.5 3.651a.651.651 0 0 1-.29.542L7.5 8L1.79 4.193A.651.651 0 0 1 2.151 3H12.85c.36 0 .651.292.651.651Zm0 2.316V11a1 1 0 0 1-1 1h-10a1 1 0 0 1-1-1V5.967c0-.2.223-.319.389-.208L7.5 9.5l5.611-3.74a.25.25 0 0 1 .389.207Z"/>`),
		g.Group(children),
	)
}