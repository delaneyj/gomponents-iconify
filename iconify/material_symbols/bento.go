package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bento(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 11V5h4q.825 0 1.413.588T22 7v4h-6ZM4 19q-.825 0-1.413-.588T2 17V7q0-.825.588-1.413T4 5h10v14H4Zm4-5.5q.625 0 1.063-.438T9.5 12q0-.625-.438-1.063T8 10.5q-.625 0-1.063.438T6.5 12q0 .625.438 1.063T8 13.5Zm8 5.5v-6h6v4q0 .825-.588 1.413T20 19h-4Z"/>`),
		g.Group(children),
	)
}