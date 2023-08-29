package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoFrameRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-.425 0-.713-.288T5 20v-1H3q-.825 0-1.413-.588T1 17V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v11q0 .825-.588 1.413T21 19h-2v1q0 .425-.288.713T18 21H6Zm.025-6H18q.3 0 .45-.275t-.05-.525l-3.5-4.675q-.15-.2-.4-.2t-.4.2L11 13.5l-2.1-2.525q-.15-.2-.388-.188T8.126 11l-2.5 3.2q-.2.25-.063.525t.463.275Z"/>`),
		g.Group(children),
	)
}