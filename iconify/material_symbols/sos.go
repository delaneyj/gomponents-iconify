package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 17q-.825 0-1.413-.588T8.5 15V9q0-.825.588-1.413T10.5 7h3q.825 0 1.413.588T15.5 9v6q0 .825-.588 1.413T13.5 17h-3ZM1 17v-2h4v-2H3q-.825 0-1.413-.588T1 11V9q0-.825.588-1.413T3 7h4v2H3v2h2q.825 0 1.413.588T7 13v2q0 .825-.588 1.413T5 17H1Zm16 0v-2h4v-2h-2q-.825 0-1.413-.588T17 11V9q0-.825.588-1.413T19 7h4v2h-4v2h2q.825 0 1.413.588T23 13v2q0 .825-.588 1.413T21 17h-4Zm-6.5-2h3V9h-3v6Z"/>`),
		g.Group(children),
	)
}