package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsMartialArts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 22l-.5-9l-3.175-1.825l-.35 1.3L8 16l-1.725 1L3.8 12.75L5 8.45l5.75-3.3L8 2.4L9.4 1L14 5.575L10.4 7.65l1.2 1.05L19.8 2L21 3.4L12.5 12L12 22h-2ZM5 7q-.825 0-1.413-.588T3 5q0-.825.588-1.413T5 3q.825 0 1.413.588T7 5q0 .825-.588 1.413T5 7Z"/>`),
		g.Group(children),
	)
}