package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarRtlThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.5 18a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM20 21.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0ZM16 23a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Zm-1.5-6.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Zm-4 1.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3ZM28 9a5 5 0 0 0-5-5H9a5 5 0 0 0-5 5v14a5 5 0 0 0 5 5h14a5 5 0 0 0 5-5V9Zm-5-3a3 3 0 0 1 3 3v1H6V9a3 3 0 0 1 3-3h14Zm3 17a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V12h20v11Z"/>`),
		g.Group(children),
	)
}