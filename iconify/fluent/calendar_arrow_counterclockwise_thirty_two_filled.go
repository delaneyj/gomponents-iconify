package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarArrowCounterclockwiseThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 9a5 5 0 0 1 5-5h14a5 5 0 0 1 5 5v1H10.292a2.5 2.5 0 0 0-4.06-.768L6 9.464v2.122l1.293-1.293a1 1 0 1 1 1.414 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 1 1 1.414-1.414L4 11.586V9Zm0 7.292V23a5 5 0 0 0 5 5h14a5 5 0 0 0 5-5V12H10.292c-.121.279-.296.54-.524.768l-3 3A2.5 2.5 0 0 1 4 16.292Zm8 .208a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM10.5 20a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm7 1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM16 15a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm7 1.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/>`),
		g.Group(children),
	)
}