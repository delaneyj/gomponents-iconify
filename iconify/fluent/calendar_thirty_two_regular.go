package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 18a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5 3.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm4 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5-6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm4 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM4 9a5 5 0 0 1 5-5h14a5 5 0 0 1 5 5v14a5 5 0 0 1-5 5H9a5 5 0 0 1-5-5V9Zm5-3a3 3 0 0 0-3 3v1h20V9a3 3 0 0 0-3-3H9ZM6 23a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V12H6v11Z"/>`),
		g.Group(children),
	)
}