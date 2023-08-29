package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateNinetyDegreesCcwOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 22q-1.275 0-2.5-.35T8.2 20.6l1.45-1.45q.775.425 1.625.638T13 20q2.925 0 4.962-2.038T20 13q0-2.925-2.038-4.963T13 6h-.15l1.55 1.55L13 9L9 5l4-4l1.4 1.45L12.85 4H13q3.75 0 6.375 2.625T22 13q0 1.875-.713 3.513t-1.924 2.85q-1.213 1.212-2.85 1.925T13 22Zm-6-3l-6-6l6-6l6 6l-6 6Zm0-2.85L10.15 13L7 9.85L3.85 13L7 16.15ZM7 13Z"/>`),
		g.Group(children),
	)
}