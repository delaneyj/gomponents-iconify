package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleAnalytics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.432 32H2.527A2.533 2.533 0 0 1 0 29.473v-5.895a2.533 2.533 0 0 1 2.527-2.527h7.577V12.63a2.54 2.54 0 0 1 2.527-2.527h8.443V2.567a2.582 2.582 0 0 1 2.567-2.568h5.792a2.583 2.583 0 0 1 2.568 2.568v26.864a2.583 2.583 0 0 1-2.568 2.568z"/>`),
		g.Group(children),
	)
}