package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMonthTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.75 3A3.25 3.25 0 0 1 25 6.25v15.5A3.25 3.25 0 0 1 21.75 25H6.25A3.25 3.25 0 0 1 3 21.75V6.25A3.25 3.25 0 0 1 6.25 3h15.5ZM8.5 15.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm5.5 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm-5.5-6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm5.5 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm5.5 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		g.Group(children),
	)
}