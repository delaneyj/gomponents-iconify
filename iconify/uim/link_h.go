package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 17H7A5 5 0 0 1 7 7h3a1 1 0 0 1 0 2H7a3 3 0 0 0 0 6h3a1 1 0 0 1 0 2zm7 0h-3a1 1 0 0 1 0-2h3a3 3 0 0 0 0-6h-3a1 1 0 0 1 0-2h3a5 5 0 0 1 0 10z"/><path fill="currentColor" d="M15 13H9a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2Z" opacity=".5"/>`),
		g.Group(children),
	)
}