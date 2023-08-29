package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GolfCourse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 21q-.625 0-1.063-.438T18 19.5q0-.625.438-1.063T19.5 18q.625 0 1.063.438T21 19.5q0 .625-.438 1.063T19.5 21ZM10 22q-2.5 0-4.25-.588T4 20q0-.575.825-1.025T7 18.25V20h2V2l8 3.9L11 9v9.05q2.15.125 3.575.663T16 20q0 .825-1.75 1.413T10 22Z"/>`),
		g.Group(children),
	)
}