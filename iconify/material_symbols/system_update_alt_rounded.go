package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemUpdateAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h4q.425 0 .713.288T9 5q0 .425-.288.713T8 6H4v12h16V6h-4q-.425 0-.713-.288T15 5q0-.425.288-.713T16 4h4q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm7-8.4V5q0-.425.288-.713T12 4q.425 0 .713.288T13 5v6.6l1.9-1.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-3.6 3.6q-.3.3-.7.3t-.7-.3l-3.6-3.6q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.9 1.9Z"/>`),
		g.Group(children),
	)
}