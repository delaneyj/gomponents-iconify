package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarRtlThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23 4a5 5 0 0 1 5 5v1H4V9a5 5 0 0 1 5-5h14Zm5 19V12H4v11a5 5 0 0 0 5 5h14a5 5 0 0 0 5-5Zm-6.5-5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM20 21.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0ZM16 23a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm-1.5-6.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm-4 1.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z"/>`),
		g.Group(children),
	)
}