package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 4a5 5 0 0 0-5 5v1h24V9a5 5 0 0 0-5-5H9ZM4 23V12h24v11a5 5 0 0 1-5 5H9a5 5 0 0 1-5-5Zm6.5-5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5 3.5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm4 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm1.5-6.5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm4 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}