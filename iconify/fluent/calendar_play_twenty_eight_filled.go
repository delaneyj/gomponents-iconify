package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarPlayTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25 9.5v5A7.5 7.5 0 0 0 14.5 25H6.25A3.25 3.25 0 0 1 3 21.75V9.5h22ZM21.75 3A3.25 3.25 0 0 1 25 6.25V8H3V6.25A3.25 3.25 0 0 1 6.25 3h15.5ZM20.5 27a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Zm-.902-9.404l4.032 2.275a.72.72 0 0 1 0 1.258l-4.032 2.274c-.49.277-1.098-.072-1.098-.629v-4.548c0-.557.609-.905 1.098-.63Z"/>`),
		g.Group(children),
	)
}