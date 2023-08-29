package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongDownC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 6.85a3.001 3.001 0 1 1 2 0l.012 12.287l1.812-1.822l1.419 1.41l-4.231 4.255l-4.255-4.23l1.41-1.42l1.845 1.835L11 6.85Zm.998-1.83a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}