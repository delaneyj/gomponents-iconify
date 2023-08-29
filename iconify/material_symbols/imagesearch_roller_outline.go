package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagesearchRollerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 23h-4q-.425 0-.713-.288T10 22v-6q0-.425.288-.713T11 15h1v-3H4q-.825 0-1.413-.588T2 10V6q0-.825.588-1.413T4 4h2V3q0-.425.288-.713T7 2h12q.425 0 .713.288T20 3v4q0 .425-.288.713T19 8H7q-.425 0-.713-.288T6 7V6H4v4h8q.825 0 1.413.588T14 12v3h1q.425 0 .713.288T16 16v6q0 .425-.288.713T15 23ZM8 4v2v-2Zm4 17h2v-4h-2v4ZM8 6h10V4H8v2Zm4 15h2h-2Z"/>`),
		g.Group(children),
	)
}