package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 12.41V14H1V5h.62l1-2h2.57l1.19 2H13v2.59l1-1V4H7L6 2H2L1 4H0v11h14v-1.59l-1-1z"/><path fill="currentColor" d="m16 8l-1-1l-2 2l-2-2l-1 1l2 2l-2 2l1 1l2-2l2 2l1-1l-2-2l2-2z"/>`),
		g.Group(children),
	)
}