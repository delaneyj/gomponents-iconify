package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayLesson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 4v7l2.5-1.5L12 11V4H7Zm11 19q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Zm-1.25-2.5l4-2.5l-4-2.5v5ZM11 17.993q0 1.082.338 2.095T12.25 22H5q-.825 0-1.413-.588T3 20V4q0-.825.588-1.413T5 2h12q.825 0 1.413.588T19 4v7.05q-.252-.025-.503-.038T17.994 11q-2.919 0-4.956 2.038Q11 15.075 11 17.993Z"/>`),
		g.Group(children),
	)
}