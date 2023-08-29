package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempleHinduOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-8q0-.425.288-.713T3 11q.425 0 .713.288T4 12v1h2l2.975-9.875V2q0-.425.288-.713T9.974 1q.425 0 .713.288t.287.712v1H13V2q0-.425.288-.713T14 1q.425 0 .713.288T15 2v1l3 10h2v-1q0-.425.288-.713T21 11q.425 0 .713.288T22 12v8q0 .825-.588 1.413T20 22h-6q-.425 0-.713-.288T13 21v-4h-2v4q0 .425-.288.713T10 22H4q-.825 0-1.413-.588T2 20Zm6.7-9h6.6l-.6-2H9.3l-.6 2Zm1.2-4h4.2l-.6-2h-3l-.6 2ZM4 20h5v-3q0-.825.588-1.413T11 15h2q.825 0 1.413.588T15 17v3h5v-5h-2q-.65 0-1.188-.388t-.737-1.037L15.9 13H8.1l-.175.575q-.2.65-.738 1.038T6 15H4v5Zm8-7Z"/>`),
		g.Group(children),
	)
}