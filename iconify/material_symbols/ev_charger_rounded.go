package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EvChargerRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.5 19l2.5-4H9.5v-3L7 16h1.5v3ZM6 10h6V5H6v5Zm7 11H5q-.425 0-.713-.288T4 20V5q0-.825.588-1.413T6 3h6q.825 0 1.413.588T14 5v7h1.25q.725 0 1.238.513T17 13.75v4.625q0 .425.35.775t.775.35q.45 0 .788-.35t.337-.775V9H19q-.425 0-.713-.288T18 8V6.5q0-.2.15-.35T18.5 6V5q0-.2.15-.35T19 4.5q.2 0 .35.15t.15.35v1h1V5q0-.2.15-.35T21 4.5q.2 0 .35.15t.15.35v1q.2 0 .35.15t.15.35V8q0 .425-.288.713T21 9h-.25v9.375q0 1.05-.763 1.838T18.126 21q-1.075 0-1.85-.788t-.775-1.837V13.75q0-.125-.063-.188t-.187-.062H14V20q0 .425-.288.713T13 21Z"/>`),
		g.Group(children),
	)
}