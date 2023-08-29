package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessCenterOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V8q0-.825.588-1.413T4 6h4V4q0-.825.588-1.413T10 2h4q.825 0 1.413.588T16 4v2h4q.825 0 1.413.588T22 8v11q0 .825-.588 1.413T20 21H4Zm6-15h4V4h-4v2Zm10 9h-5v1q0 .425-.288.713T14 17h-4q-.425 0-.713-.288T9 16v-1H4v4h16v-4Zm-9 0h2v-2h-2v2Zm-7-2h5v-1q0-.425.288-.713T10 11h4q.425 0 .713.288T15 12v1h5V8H4v5Zm8 1Z"/>`),
		g.Group(children),
	)
}