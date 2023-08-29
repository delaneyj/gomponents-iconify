package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.85 22.5l-1.4-1.4l1.1-1.1H5v-2h4.55l-1.1-1.1l1.4-1.4l3.5 3.5l-3.5 3.5ZM14 20v-2h5v2h-5Zm-9-3V4q0-.825.588-1.413T7 2h10q.825 0 1.413.588T19 4v13h-6.225L9.85 14.075L6.925 17H5Zm2-3.85q1.2-.575 2.463-.863T12 12q1.275 0 2.538.288T17 13.15V4H7v9.15ZM12 11q-1.25 0-2.125-.875T9 8q0-1.25.875-2.125T12 5q1.25 0 2.125.875T15 8q0 1.25-.875 2.125T12 11Z"/>`),
		g.Group(children),
	)
}