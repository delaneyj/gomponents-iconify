package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarArrowCounterclockwiseTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 17.75V8.5H7.483a1.747 1.747 0 0 1-.495 1.487l-2 2A1.75 1.75 0 0 1 3 12.331v5.419A3.25 3.25 0 0 0 6.25 21h11.5A3.25 3.25 0 0 0 21 17.75Zm-15-1.5a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Zm4.75 0a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0ZM6 11.75a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Zm4.75 0a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Zm4.75 0a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0ZM21 7v-.75A3.25 3.25 0 0 0 17.75 3H6.25A3.25 3.25 0 0 0 3 6.25v2.69l-.72-.72a.75.75 0 0 0-1.06 1.06l2 2a.75.75 0 0 0 1.06 0l2-2a.75.75 0 0 0-1.06-1.06l-.72.72V7H21Z"/>`),
		g.Group(children),
	)
}