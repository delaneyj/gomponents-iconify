package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeySkeleton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 4.41l.71-.7a1 1 0 1 0-1.42-1.42l-1.4 1.41l-2.83 2.83l-6.31 6.3a5 5 0 1 0 1.42 1.42l5.59-5.6l2.12 2.13a1 1 0 1 0 1.41-1.42l-2.12-2.12l1.42-1.41l.7.7a1 1 0 1 0 1.42-1.41ZM7 20a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/>`),
		g.Group(children),
	)
}