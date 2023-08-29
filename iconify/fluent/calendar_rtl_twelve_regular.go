package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarRtlTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 5.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0ZM8.5 7a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1ZM7 5.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0ZM6.5 7a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1ZM5 5.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Zm6-2A2.5 2.5 0 0 0 8.5 1h-5A2.5 2.5 0 0 0 1 3.5v5A2.5 2.5 0 0 0 3.5 11h5A2.5 2.5 0 0 0 11 8.5v-5ZM8.5 2a1.5 1.5 0 0 1 1.415 1h-7.83A1.5 1.5 0 0 1 3.5 2h5ZM2 4h8v4.5A1.5 1.5 0 0 1 8.5 10h-5A1.5 1.5 0 0 1 2 8.5V4Z"/>`),
		g.Group(children),
	)
}