package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempleHinduRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.8 7l1.175-3.875V2q0-.425.288-.713T9.974 1q.425 0 .713.288t.287.712v1H13V2q0-.425.288-.713T14 1q.425 0 .713.288T15 2v1l1.2 4H7.8ZM2 20v-8q0-.425.288-.713T3 11q.425 0 .713.288T4 12v1h16v-1q0-.425.288-.713T21 11q.425 0 .713.288T22 12v8q0 .825-.588 1.413T20 22h-6q-.425 0-.713-.288T13 21v-3q0-.425-.288-.713T12 17q-.425 0-.713.288T11 18v3q0 .425-.288.713T10 22H4q-.825 0-1.413-.588T2 20Zm4.6-9l.6-2h9.6l.6 2H6.6Z"/>`),
		g.Group(children),
	)
}