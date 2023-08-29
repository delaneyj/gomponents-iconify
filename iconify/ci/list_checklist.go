package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListChecklist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 17h9M8 15l-2.5 3L4 17m7-5h9M8 10l-2.5 3L4 12m7-5h9M8 5L5.5 8L4 7"/>`),
		g.Group(children),
	)
}