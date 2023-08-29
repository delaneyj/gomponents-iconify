package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 22.5l-1.6-1.2L7 17.825V12.5q0-.775.138-1.713T7.5 9.1L6 9.95v3.55H4V8.8l5.4-3.075q.2-.125.425-.175t.475-.05q.6 0 1.1.3t.75.825l.775 1.675q.5 1.1 1.525 1.65t2.55.55v2h-.975l5.475 9.55l-.875.5L14.7 12.225q-1-.325-1.813-.937T11.5 9.8q-.25.725-.388 1.663t-.087 1.737L13 16v6.5h-2v-5l-1.775-2.55L9 18.5l-3 4ZM11.5 5q-.825 0-1.413-.588T9.5 3q0-.825.588-1.413T11.5 1q.825 0 1.413.588T13.5 3q0 .825-.588 1.413T11.5 5Z"/>`),
		g.Group(children),
	)
}