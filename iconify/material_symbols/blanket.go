package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blanket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-1.25 0-2.125-.875T1 18V6q0-1.25.875-2.125T4 3h8q1.25 0 2.125.875T15 6v6.3h1q.825 0 1.413.588T18 14.3v3.2q0 .2.15.35t.35.15q.2 0 .35-.15t.15-.35V11q-.825 0-1.413-.588T17 9V6h1V4h1.5v2h1V4H22v2h1v3q0 .825-.588 1.413T21 11v6.4q0 1.05-.725 1.775T18.5 19.9q-1.05 0-1.775-.725T16 17.4v-3.2h-1V19q0-1.575-1.213-2.788T11 15H6.5q-1.05 0-1.775.725T4 17.5q0 1.05.725 1.775T6.5 20H11q.425 0 .713-.288T12 19q0-.425-.288-.713T11 18H6.5q-.2 0-.35-.15T6 17.5q0-.2.15-.35T6.5 17H11q.825 0 1.413.588T13 19q0 .825-.588 1.413T11 21H4Z"/>`),
		g.Group(children),
	)
}