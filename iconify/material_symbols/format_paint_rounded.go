package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatPaintRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22q-.825 0-1.413-.588T9 20v-4H6q-.825 0-1.413-.588T4 14V7q0-1.65 1.175-2.825T8 3h12v11q0 .825-.588 1.413T18 16h-3v4q0 .825-.588 1.413T13 22h-2ZM6 10h12V5h-1v3q0 .425-.288.713T16 9q-.425 0-.713-.288T15 8V5h-1v1q0 .425-.288.713T13 7q-.425 0-.713-.288T12 6V5H8q-.825 0-1.413.588T6 7v3Z"/>`),
		g.Group(children),
	)
}