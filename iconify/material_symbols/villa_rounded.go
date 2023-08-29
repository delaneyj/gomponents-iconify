package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VillaRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21v-7q0-.825.588-1.413T11 12h6q0-.825.588-1.413T19 10q.825 0 1.413.588T21 12v9h-5v-4q0-.425-.288-.713T15 16q-.425 0-.713.288T14 17v4H9Zm-6 0V9.375q0-.625.35-1.137t.925-.738L14.65 3.525q.5-.2.925.1T16 4.45V10h-6q-1.25 0-2.125.875T7 13v8H3Z"/>`),
		g.Group(children),
	)
}