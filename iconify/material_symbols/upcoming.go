package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upcoming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19v-5q0-.825.588-1.413T4 12h5q0 1.25.875 2.125T12 15q1.25 0 2.125-.875T15 12h5q.825 0 1.413.588T22 14v5q0 .825-.588 1.413T20 21H4Zm13.6-10.2l-1.4-1.4l3.55-3.55l1.4 1.4l-3.55 3.55Zm-11.2 0L2.85 7.25l1.4-1.4L7.8 9.4l-1.4 1.4ZM11 8V3h2v5h-2Z"/>`),
		g.Group(children),
	)
}