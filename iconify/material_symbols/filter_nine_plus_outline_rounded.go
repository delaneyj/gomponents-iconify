package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterNinePlusOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14q.825 0 1.413-.588T14 12V8q0-.825-.588-1.413T12 6h-1q-.825 0-1.413.588T9 8v1q0 .825.588 1.413T11 11h1v1h-1q-.425 0-.713.288T10 13q0 .425.288.713T11 14h1Zm0-5h-1V8h1v1Zm8-5Zm-3.5 7v1q0 .425.288.713T17.5 13q.425 0 .713-.288T18.5 12v-1H20V9h-1.5V8q0-.425-.288-.713T17.5 7q-.425 0-.713.288T16.5 8v1h-1q-.425 0-.713.288T14.5 10q0 .425.288.713T15.5 11h1ZM8 18q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm0-2h12V4H8v12Zm-4 6q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4ZM8 4v12V4Z"/>`),
		g.Group(children),
	)
}