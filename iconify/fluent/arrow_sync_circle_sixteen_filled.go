package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSyncCircleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 2a6 6 0 1 0 0 12A6 6 0 0 0 8 2ZM5.5 8H7l.09.008A.5.5 0 0 1 7 9h-.333a1.667 1.667 0 0 0 2.533.156a.5.5 0 1 1 .72.694A2.665 2.665 0 0 1 6 9.764v.237l-.008.09a.5.5 0 0 1-.992-.09V8.5l.008-.09A.5.5 0 0 1 5.5 8ZM8 5.334a2.66 2.66 0 0 1 2 .903v-.234l.008-.09a.5.5 0 0 1 .992.09V7.5l-.008.09A.5.5 0 0 1 10.5 8H9.004l-.09-.008A.5.5 0 0 1 9.004 7h.329a1.667 1.667 0 0 0-2.537-.152a.5.5 0 0 1-.723-.691A2.66 2.66 0 0 1 8 5.334Z"/>`),
		g.Group(children),
	)
}