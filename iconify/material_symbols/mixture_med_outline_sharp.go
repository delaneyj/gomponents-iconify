package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MixtureMedOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22.5L6 21v-4H5q-.825 0-1.413-.588T3 15V7.5H2v-2h4V4H4.5V2h5v2H8v1.5h4v2h-1V15q0 .825-.588 1.413T9 17H8v5.5Zm9-.5q-1.65 0-2.825-1.175T13 18v-8q0-1.65 1.175-2.825T17 6q1.65 0 2.825 1.175T21 10v8q0 1.65-1.175 2.825T17 22ZM5 15h4v-1.5H6.5V12H9v-1.5H6.5V9H9V7.5H5V15Zm12-7q-.825 0-1.413.588T15 10h4q0-.825-.588-1.413T17 8Zm-2 8h4v-4h-4v4Zm2 4q.825 0 1.413-.588T19 18h-4q0 .825.588 1.413T17 20Zm0-10Zm0 8Z"/>`),
		g.Group(children),
	)
}