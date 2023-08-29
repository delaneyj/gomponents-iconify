package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarThreeWeekFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 0a2 2 0 0 0-2 2h16a2 2 0 0 0-2-2H2zM0 14V3h16v11a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2zm12-8a1 1 0 1 0 2 0a1 1 0 0 0-2 0zM5 9a1 1 0 1 0 2 0a1 1 0 0 0-2 0zm5-2a1 1 0 1 1 0-2a1 1 0 0 1 0 2zM2 9a1 1 0 1 0 2 0a1 1 0 0 0-2 0z"/>`),
		g.Group(children),
	)
}