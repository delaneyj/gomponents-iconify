package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentControlButtonMoveCircleMoveButtonCircleDirectionArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M5.5 4L7 3l1.5 1m-3 6L7 11l1.5-1M10 5.5L11 7l-1 1.5m-6-3L3 7l1 1.5"/></g>`),
		g.Group(children),
	)
}