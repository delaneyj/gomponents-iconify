package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oncology(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.575 23l-2.55-2.55q-.45.275-.95.413T17 21q-1.65 0-2.825-1.175T13 17q0-1.65 1.175-2.825T17 13q1.65 0 2.825 1.175T21 17q0 .575-.15 1.088t-.425.962l2.55 2.55l-1.4 1.4ZM17 19q.825 0 1.413-.588T19 17q0-.825-.588-1.413T17 15q-.825 0-1.413.588T15 17q0 .825.588 1.413T17 19ZM3 22v-6q0-1.25.875-2.125T6 13h2q1.25 0 2.125-.875T11 10q0-.425-.288-.713T10 9q-.825 0-1.413-.588T8 7V2h6v2q2.925 0 4.963 2.038T21 11v1.525q-.825-.725-1.85-1.125T17 11q-2.5 0-4.25 1.75T11 17q0 .525.088 1.025t.262.975H10q-.425 0-.713.288T9 20v2H3Z"/>`),
		g.Group(children),
	)
}