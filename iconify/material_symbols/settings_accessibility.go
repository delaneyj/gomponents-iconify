package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsAccessibility(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4q-.825 0-1.413-.588T10 2q0-.825.588-1.413T12 0q.825 0 1.413.588T14 2q0 .825-.588 1.413T12 4ZM9 19V7q-1.5-.125-3.05-.375T3 6l.5-2q1.95.525 4.15.763T12 5q2.15 0 4.35-.238T20.5 4l.5 2q-1.4.375-2.95.625T15 7v12h-2v-6h-2v6H9Zm-1 5q-.425 0-.713-.288T7 23q0-.425.288-.713T8 22q.425 0 .713.288T9 23q0 .425-.288.713T8 24Zm4 0q-.425 0-.713-.288T11 23q0-.425.288-.713T12 22q.425 0 .713.288T13 23q0 .425-.288.713T12 24Zm4 0q-.425 0-.713-.288T15 23q0-.425.288-.713T16 22q.425 0 .713.288T17 23q0 .425-.288.713T16 24Z"/>`),
		g.Group(children),
	)
}