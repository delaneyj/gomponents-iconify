package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19q-.825 0-1.413-.588T2 17V7q0-.825.588-1.413T4 5h16q.825 0 1.413.588T22 7v10q0 .825-.588 1.413T20 19H4Zm5-3h6q.425 0 .713-.288T16 15q0-.425-.288-.713T15 14H9q-.425 0-.713.288T8 15q0 .425.288.713T9 16Zm-4-3h2v-2H5v2Zm3 0h2v-2H8v2Zm3 0h2v-2h-2v2Zm3 0h2v-2h-2v2Zm3 0h2v-2h-2v2ZM5 10h2V8H5v2Zm3 0h2V8H8v2Zm3 0h2V8h-2v2Zm3 0h2V8h-2v2Zm3 0h2V8h-2v2Z"/>`),
		g.Group(children),
	)
}