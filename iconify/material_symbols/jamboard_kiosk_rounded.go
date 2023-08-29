package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JamboardKioskRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21q-.425 0-.713-.288T6 20q0-.425.288-.713T7 19h4v-3H4q-.825 0-1.413-.588T2 14V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v9q0 .825-.588 1.413T20 16h-7v3h4q.425 0 .713.288T18 20q0 .425-.288.713T17 21H7Zm-3-7h16V5H4v9Z"/>`),
		g.Group(children),
	)
}