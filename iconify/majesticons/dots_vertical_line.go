package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsVerticalLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9 5a3 3 0 1 0 6 0a3 3 0 0 0-6 0zm3 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm-1-3a1 1 0 1 0 2 0a1 1 0 0 0-2 0zm1 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6zm-1-3a1 1 0 1 0 2 0a1 1 0 0 0-2 0z"/></g>`),
		g.Group(children),
	)
}