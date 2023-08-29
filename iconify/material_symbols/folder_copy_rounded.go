package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderCopyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V7q0-.425.288-.713T2 6q.425 0 .713.288T3 7v12h16q.425 0 .713.288T20 20q0 .425-.288.713T19 21H3Zm4-4q-.825 0-1.413-.588T5 15V4q0-.825.588-1.413T7 2h4.175q.4 0 .763.15t.637.425L14 4h7q.825 0 1.413.588T23 6v9q0 .825-.588 1.413T21 17H7Z"/>`),
		g.Group(children),
	)
}