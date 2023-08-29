package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Destruction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20v-4q0-.825.588-1.413T5 14h14q.825 0 1.413.588T21 16v4q0 .825-.588 1.413T19 22H5Zm1.525-9.5L1.5 9.65l6-1l-1.625-5.875L10.85 6.3l3-5.3l1 6l5.875-1.625L17.2 10.35L21 12.5H6.525Z"/>`),
		g.Group(children),
	)
}