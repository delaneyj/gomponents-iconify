package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMonthThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 7.5A4.5 4.5 0 0 1 7.5 3h17A4.5 4.5 0 0 1 29 7.5v17a4.5 4.5 0 0 1-4.5 4.5h-17A4.5 4.5 0 0 1 3 24.5v-17Zm9 5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0ZM10.5 20a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm5.5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-6a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm5.5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}