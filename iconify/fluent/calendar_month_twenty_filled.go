package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMonthTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 3A2.5 2.5 0 0 1 17 5.5v9a2.5 2.5 0 0 1-2.5 2.5h-9A2.5 2.5 0 0 1 3 14.5v-9A2.5 2.5 0 0 1 5.5 3h9ZM7 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM7 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}