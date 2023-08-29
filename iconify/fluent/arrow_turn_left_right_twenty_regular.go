package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnLeftRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.996 14.566a.5.5 0 0 1-.31.4l-5 1.998a.5.5 0 0 1-.371-.928l3.837-1.535l-11.524-4.59c-.828-.329-.843-1.495-.024-1.847l11.7-5.023a.5.5 0 0 1 .394.919l-11.7 5.023l11.468 4.567l-1.913-3.826a.5.5 0 0 1 .894-.447l2.491 4.982a.5.5 0 0 1 .058.307Z"/>`),
		g.Group(children),
	)
}