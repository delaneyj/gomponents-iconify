package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarClockTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 6.25A3.25 3.25 0 0 0 17.75 3H6.25A3.25 3.25 0 0 0 3 6.25v11.5A3.25 3.25 0 0 0 6.25 21h5.772a6.471 6.471 0 0 1-.709-1.5H6.25a1.75 1.75 0 0 1-1.75-1.75V8.5h15v2.813a6.471 6.471 0 0 1 1.5.709V6.25ZM6.25 4.5h11.5c.966 0 1.75.784 1.75 1.75V7h-15v-.75c0-.966.784-1.75 1.75-1.75ZM23 17.5a5.5 5.5 0 1 0-11 0a5.5 5.5 0 0 0 11 0Zm-5.5 0h2a.5.5 0 0 1 0 1H17a.5.5 0 0 1-.5-.491v-3.01a.5.5 0 0 1 1 0V17.5Z"/>`),
		g.Group(children),
	)
}