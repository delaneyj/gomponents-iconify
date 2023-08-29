package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SurveillanceCamerasTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19.006 26.276V37H5m37.62-15.785l-3.863-1.035l-4.003 7.21l5.796 1.553l2.07-7.728Z"/><path d="m38.757 20.18l-4.003 7.21l-1.742 2.639L5 22.523L8.623 9L40.5 17.541l-1.742 2.639Z"/></g>`),
		g.Group(children),
	)
}