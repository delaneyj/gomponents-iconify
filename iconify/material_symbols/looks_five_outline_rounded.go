package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LooksFiveOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 17q.825 0 1.413-.588T15 15v-2q0-.825-.588-1.413T13 11h-2V9h3q.425 0 .713-.288T15 8q0-.425-.288-.713T14 7h-4q-.425 0-.713.288T9 8v4q0 .425.288.713T10 13h3v2h-3q-.425 0-.713.288T9 16q0 .425.288.713T10 17h3Zm-8 4q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}