package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUndoSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.75a.75.75 0 0 1 1.5 0v3.095l2.673-2.673a4 4 0 0 1 5.657 5.656l-4.95 4.95a.75.75 0 1 1-1.06-1.06l4.95-4.95a2.5 2.5 0 0 0-3.536-3.536L5.966 6.5H8.25a.75.75 0 0 1 0 1.5h-4.4A.85.85 0 0 1 3 7.15v-4.4Z"/>`),
		g.Group(children),
	)
}