package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluidBalance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6h-7q-1.425 0-2.488.85T11.125 9H6v2h5v2H6v2h5q0 1.475.537 2.763T13.075 20H4Zm15 3q-.825 0-1.413-.588T17 21v-1.1q-1.725-.35-2.863-1.713T13 15v-5q0-.825.575-1.413T15 8h6q.825 0 1.413.588T23 10v5q0 1.825-1.137 3.188T19 19.9V21h3v2h-3Zm.75-9H21v-4h-6v2h.75q.825 0 1.563.375T18.55 13.4q.2.3.525.45t.675.15Z"/>`),
		g.Group(children),
	)
}