package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminalRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V8H4v10Zm4.675-5l-1.9-1.9q-.3-.3-.288-.7t.313-.7q.3-.275.7-.287t.7.287l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.687.288T6.8 16.3q-.275-.275-.275-.7t.275-.7L8.675 13ZM13 17q-.425 0-.713-.288T12 16q0-.425.288-.713T13 15h4q.425 0 .713.288T18 16q0 .425-.288.713T17 17h-4Z"/>`),
		g.Group(children),
	)
}