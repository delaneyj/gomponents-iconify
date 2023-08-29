package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowMaximize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 8h20v2H2z" opacity=".25"/><path fill="currentColor" d="M3 2h18a1 1 0 0 1 1 1v5H2V3a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M2 10h20v11a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V10z" opacity=".5"/>`),
		g.Group(children),
	)
}