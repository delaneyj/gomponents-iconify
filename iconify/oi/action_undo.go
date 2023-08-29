package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActionUndo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4.5 1C2.57 1 1 2.57 1 4.5V5H0l2 2l2-2H3v-.5a2.5 2.5 0 0 1 5 0C8 2.57 6.43 1 4.5 1z"/>`),
		g.Group(children),
	)
}