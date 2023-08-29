package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DuplicateLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16 6a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v6a4 4 0 0 0 4 4h2v2a4 4 0 0 0 4 4h6a4 4 0 0 0 4-4v-6a4 4 0 0 0-4-4h-2V6zm-2 2h-2a4 4 0 0 0-4 4v2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2zm4 2a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-6a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h6z"/></g>`),
		g.Group(children),
	)
}