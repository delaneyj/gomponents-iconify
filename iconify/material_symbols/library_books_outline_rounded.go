package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LibraryBooksOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 14h2q.425 0 .713-.288T14 13q0-.425-.288-.713T13 12h-2q-.425 0-.713.288T10 13q0 .425.288.713T11 14Zm0-3h6q.425 0 .713-.288T18 10q0-.425-.288-.713T17 9h-6q-.425 0-.713.288T10 10q0 .425.288.713T11 11Zm0-3h6q.425 0 .713-.288T18 7q0-.425-.288-.713T17 6h-6q-.425 0-.713.288T10 7q0 .425.288.713T11 8ZM8 18q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm0-2h12V4H8v12Zm-4 6q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4ZM8 4v12V4Z"/>`),
		g.Group(children),
	)
}