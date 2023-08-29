package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 22v-7.5H4V9q0-.825.588-1.413T6 7h3q.825 0 1.413.588T11 9v5.5H9.5V22h-4Zm2-16q-.825 0-1.413-.588T5.5 4q0-.825.588-1.413T7.5 2q.825 0 1.413.588T9.5 4q0 .825-.588 1.413T7.5 6ZM15 22v-6h-3l2.55-7.65q.2-.65.738-1T16.5 7q.675 0 1.213.35t.737 1L21 16h-3v6h-3Zm1.5-16q-.825 0-1.413-.588T14.5 4q0-.825.588-1.413T16.5 2q.825 0 1.413.588T18.5 4q0 .825-.588 1.413T16.5 6Z"/>`),
		g.Group(children),
	)
}