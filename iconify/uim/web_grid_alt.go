package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebGridAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 8H2v2h9v12h2V10h9z" opacity=".25"/><path fill="currentColor" d="M3 2h18a1 1 0 0 1 1 1v5H2V3a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M2 10h9v12H3a1 1 0 0 1-1-1V10zm11 0h9v11a1 1 0 0 1-1 1h-8V10z" opacity=".5"/>`),
		g.Group(children),
	)
}