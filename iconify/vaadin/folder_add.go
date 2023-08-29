package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 6V4H7L6 2H2L1 4H0v11h14v-1H1V5h.62l1-2h2.57l1.19 2H13v1h1z"/><path fill="currentColor" d="M14 7h-2v2h-2v2h2v2h2v-2h2V9h-2V7z"/>`),
		g.Group(children),
	)
}