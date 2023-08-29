package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolunteerActivismRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 20v-7q0-.825.588-1.413T3 11q.825 0 1.413.588T5 13v7q0 .825-.588 1.413T3 22q-.825 0-1.413-.588T1 20Zm12.375 1.825L7 20.025V11h1.6q.175 0 .35.038t.35.087l6.925 2.575q.35.125.563.45t.212.675q0 .525-.363.85T15.8 16h-2.625q-.125 0-.188-.013t-.162-.062L11.7 15.5q-.2-.075-.4.025t-.25.275q-.05.2.025.375t.275.25l1.475.525q.05.025.15.038t.175.012H20q.8 0 1.4.575T22 19l-7.375 2.775q-.25.1-.613.113t-.637-.063ZM16 12.675q-.175 0-.338-.05t-.287-.175q-2.15-1.95-3.763-3.763T10 5.3q0-1.4.95-2.35T13.3 2q.8 0 1.5.338t1.2.912q.5-.575 1.2-.912T18.7 2q1.4 0 2.35.95T22 5.3q0 1.575-1.613 3.388t-3.762 3.762q-.125.125-.288.175t-.337.05Z"/>`),
		g.Group(children),
	)
}