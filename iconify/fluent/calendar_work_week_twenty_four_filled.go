package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarWorkWeekTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 11V8.5h7V11h-7Zm9.25 10A3.25 3.25 0 0 0 21 17.75V6.25A3.25 3.25 0 0 0 17.75 3H6.25A3.25 3.25 0 0 0 3 6.25v11.5A3.25 3.25 0 0 0 6.25 21h11.5Zm-10-14h8.5a.75.75 0 0 1 .75.75v4a.75.75 0 0 1-.75.75h-8.5a.75.75 0 0 1-.75-.75v-4A.75.75 0 0 1 7.75 7Z"/>`),
		g.Group(children),
	)
}