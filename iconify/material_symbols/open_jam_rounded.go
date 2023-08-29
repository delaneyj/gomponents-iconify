package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenJamRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21q-.425 0-.713-.288T6 20q0-.425.288-.713T7 19h6v-7.2l.9.9q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7l-2.6-2.6q-.3-.3-.7-.3t-.7.3l-2.6 2.6q-.275.275-.275.7t.275.7q.275.275.7.275t.7-.275l.9-.9V16H4q-.825 0-1.413-.588T2 14V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v9q0 .825-.588 1.413T20 16h-5v3h2q.425 0 .713.288T18 20q0 .425-.288.713T17 21H7Z"/>`),
		g.Group(children),
	)
}