package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterThreeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 15q.825 0 1.413-.588T17 13v-1.5q0-.65-.425-1.075T15.5 10q.65 0 1.075-.425T17 8.5V7q0-.825-.588-1.413T15 5h-3q-.425 0-.713.288T11 6q0 .425.288.713T12 7h3v2h-1q-.425 0-.713.288T13 10q0 .425.288.713T14 11h1v2h-3q-.425 0-.713.288T11 14q0 .425.288.713T12 15h3Zm-7 3q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm0-2h12V4H8v12Zm-4 6q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4ZM8 4v12V4Z"/>`),
		g.Group(children),
	)
}