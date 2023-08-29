package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WcSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 22v-7.5H4V7h7v7.5H9.5V22h-4Zm9.5 0v-6h-3l3-9h3l3 9h-3v6h-3ZM7.5 6q-.825 0-1.413-.588T5.5 4q0-.825.588-1.413T7.5 2q.825 0 1.413.588T9.5 4q0 .825-.588 1.413T7.5 6Zm9 0q-.825 0-1.413-.588T14.5 4q0-.825.588-1.413T16.5 2q.825 0 1.413.588T18.5 4q0 .825-.588 1.413T16.5 6Z"/>`),
		g.Group(children),
	)
}