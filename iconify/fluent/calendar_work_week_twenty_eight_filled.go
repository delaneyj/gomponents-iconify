package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarWorkWeekTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 13.5v-4h10v4H9Zm-6 8.25A3.25 3.25 0 0 0 6.25 25h15.5A3.25 3.25 0 0 0 25 21.75V6.25A3.25 3.25 0 0 0 21.75 3H6.25A3.25 3.25 0 0 0 3 6.25v15.5ZM8.25 8h11.5a.75.75 0 0 1 .75.75v5.5a.75.75 0 0 1-.75.75H8.25a.75.75 0 0 1-.75-.75v-5.5A.75.75 0 0 1 8.25 8Z"/>`),
		g.Group(children),
	)
}