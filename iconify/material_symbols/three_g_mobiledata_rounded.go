package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeGMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17q-.425 0-.713-.288T3 16q0-.425.288-.713T4 15h4v-2H4q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h4V9H4q-.425 0-.713-.288T3 8q0-.425.288-.713T4 7h4q.825 0 1.413.588T10 9v1.5q0 .625-.438 1.063T8.5 12q.625 0 1.063.438T10 13.5V15q0 .825-.588 1.413T8 17H4Zm10 0q-.825 0-1.413-.588T12 15V9q0-.825.588-1.413T14 7h6q.425 0 .713.288T21 8q0 .425-.288.713T20 9h-6v6h5v-2h-1.5q-.425 0-.713-.288T16.5 12q0-.425.288-.713T17.5 11H20q.425 0 .713.288T21 12v3q0 .825-.588 1.413T19 17h-5Z"/>`),
		g.Group(children),
	)
}