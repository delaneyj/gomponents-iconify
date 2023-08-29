package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsHorizontalLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 12a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm3-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm4 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm3-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm7-2a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm-1 3a1 1 0 1 1 2 0a1 1 0 0 1-2 0z"/></g>`),
		g.Group(children),
	)
}