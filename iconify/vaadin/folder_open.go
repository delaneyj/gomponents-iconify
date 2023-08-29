package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 6V4H7L6 2H2L1 4H0v9.5L3 6z"/><path fill="currentColor" d="M3.7 7L.5 15h12.8l2.5-8z"/>`),
		g.Group(children),
	)
}