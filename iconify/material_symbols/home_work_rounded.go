package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeWorkRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21v-9.975q0-.5-.213-.925t-.612-.7L10 4.95q0-.8.588-1.375T12 3h9q.825 0 1.413.588T23 5v14q0 .825-.588 1.413T21 21h-4Zm0-4h2v-2h-2v2Zm0-4h2v-2h-2v2Zm0-4h2V7h-2v2ZM1 20v-7.975q0-.5.225-.925t.625-.7l5-3.575Q7.375 6.45 8 6.45t1.15.375l5 3.575q.4.275.625.7t.225.925V20q0 .425-.288.713T14 21h-4v-6H6v6H2q-.425 0-.713-.288T1 20Z"/>`),
		g.Group(children),
	)
}