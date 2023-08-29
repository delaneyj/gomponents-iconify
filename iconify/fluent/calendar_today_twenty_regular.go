package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarTodayTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 3A2.5 2.5 0 0 1 17 5.5v9a2.5 2.5 0 0 1-2.5 2.5h-3v-1h3a1.5 1.5 0 0 0 1.5-1.5V7H4v7.5A1.5 1.5 0 0 0 5.5 16h3v1h-3A2.5 2.5 0 0 1 3 14.5v-9A2.5 2.5 0 0 1 5.5 3h9Zm0 1h-9A1.5 1.5 0 0 0 4 5.5V6h12v-.5A1.5 1.5 0 0 0 14.5 4ZM11 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm.884 5.07a.5.5 0 0 1-.704.064l-.68-.566V17.5a.5.5 0 1 1-1 0v-3.932l-.68.566a.5.5 0 0 1-.64-.768l1.5-1.25a.5.5 0 0 1 .64 0l1.5 1.25a.5.5 0 0 1 .064.704Z"/>`),
		g.Group(children),
	)
}