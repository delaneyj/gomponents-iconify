package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarRtlTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.95 3A2.5 2.5 0 0 0 8.5 1h-5a2.5 2.5 0 0 0-2.45 2h9.9ZM11 4v4.5A2.5 2.5 0 0 1 8.5 11h-5A2.5 2.5 0 0 1 1 8.5V4h10ZM9 5.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Zm0 2a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0ZM6.5 5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM7 7.5a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0ZM4.5 5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z"/>`),
		g.Group(children),
	)
}