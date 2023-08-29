package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 13.47V14H1V5h.62l1-2h2.57l1.19 2H13v.91c.385.179.716.407 1.001.681L14 4H7L6 2H2L1 4H0v11h14v-.53z"/><path fill="currentColor" d="m15.78 12.72l-1.92-1.92a.727.727 0 0 0-.325-.179a3.014 3.014 0 0 0 .468-1.618a3 3 0 1 0-1.371 2.52c.02.136.083.248.169.337l1.92 1.92a.75.75 0 0 0 1.059-1.061zM11 11a2 2 0 1 1-.001-3.999A2 2 0 0 1 11 11z"/>`),
		g.Group(children),
	)
}