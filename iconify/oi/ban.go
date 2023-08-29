package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ban(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 0C1.8 0 0 1.8 0 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4zm0 1c.66 0 1.26.21 1.75.56L1.56 5.75C1.21 5.26 1 4.66 1 4c0-1.66 1.34-3 3-3zm2.44 1.25C6.79 2.74 7 3.34 7 4c0 1.66-1.34 3-3 3c-.66 0-1.26-.21-1.75-.56l4.19-4.19z"/>`),
		g.Group(children),
	)
}