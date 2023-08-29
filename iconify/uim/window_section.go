package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowSection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 10h6v12H9zm-7 0v11a1 1 0 0 0 1 1h4V10H2z" opacity=".5"/><path fill="currentColor" d="M22 8H2v2h5v12h2V10h6v12h2V10h5z" opacity=".25"/><path fill="currentColor" d="M17 10v12h4a1 1 0 0 0 1-1V10h-5z" opacity=".5"/><path fill="currentColor" d="M3 2h18a1 1 0 0 1 1 1v5H2V3a1 1 0 0 1 1-1z"/>`),
		g.Group(children),
	)
}