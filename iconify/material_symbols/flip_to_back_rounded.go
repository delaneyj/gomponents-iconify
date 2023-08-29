package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipToBackRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9V7h2v2H7Zm0 4v-2h2v2H7Zm0-8q0-.825.588-1.413T9 3v2H7Zm4 12v-2h2v2h-2Zm8-12V3q.825 0 1.413.588T21 5h-2Zm-8 0V3h2v2h-2ZM9 17q-.825 0-1.413-.588T7 15h2v2Zm10-4v-2h2v2h-2Zm0-4V7h2v2h-2Zm0 8v-2h2q0 .825-.588 1.413T19 17ZM5 21q-.825 0-1.413-.588T3 19V8q0-.425.288-.713T4 7q.425 0 .713.288T5 8v11h11q.425 0 .713.288T17 20q0 .425-.288.713T16 21H5ZM15 5V3h2v2h-2Zm0 12v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}