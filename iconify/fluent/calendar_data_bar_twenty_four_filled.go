package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarDataBarTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 8.5v4.55a2.49 2.49 0 0 0-1 .45a2.5 2.5 0 0 0-5 0v2a2.5 2.5 0 0 0-4 2V21H6.25A3.25 3.25 0 0 1 3 17.75V8.5h18ZM17.75 3A3.25 3.25 0 0 1 21 6.25V7H3v-.75A3.25 3.25 0 0 1 6.25 3h11.5Zm-.25 9a1.5 1.5 0 0 0-1.5 1.5v8a1.5 1.5 0 0 0 3 0v-8a1.5 1.5 0 0 0-1.5-1.5Zm-4 4a1.5 1.5 0 0 0-1.5 1.5v4a1.5 1.5 0 0 0 3 0v-4a1.5 1.5 0 0 0-1.5-1.5Zm6.5-.5a1.5 1.5 0 0 1 3 0v6a1.5 1.5 0 0 1-3 0v-6Z"/>`),
		g.Group(children),
	)
}