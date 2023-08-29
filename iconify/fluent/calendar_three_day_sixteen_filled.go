package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThreeDaySixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 11.5a2.5 2.5 0 0 1-2.5 2.5h-7A2.5 2.5 0 0 1 2 11.5v-7A2.5 2.5 0 0 1 4.5 2h7A2.5 2.5 0 0 1 14 4.5v7Zm-8-6a.5.5 0 0 0-1 0v5a.5.5 0 0 0 1 0v-5Zm2.5 0a.5.5 0 0 0-1 0v5a.5.5 0 0 0 1 0v-5Zm2.5 0a.5.5 0 0 0-1 0v5a.5.5 0 0 0 1 0v-5Z"/>`),
		g.Group(children),
	)
}