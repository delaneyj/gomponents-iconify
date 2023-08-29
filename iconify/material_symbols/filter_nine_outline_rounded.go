package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterNineOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 15q.825 0 1.413-.588T17 13V7q0-.825-.588-1.413T15 5h-2q-.825 0-1.413.588T11 7v2q0 .825.588 1.413T13 11h2v2h-2q-.425 0-.713.288T12 14q0 .425.288.713T13 15h2Zm0-6h-2V7h2v2Zm-7 9q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm0-2h12V4H8v12Zm-4 6q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4ZM8 4v12V4Z"/>`),
		g.Group(children),
	)
}