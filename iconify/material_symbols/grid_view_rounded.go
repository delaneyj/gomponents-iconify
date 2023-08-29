package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridViewRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 11q-.825 0-1.413-.588T3 9V5q0-.825.588-1.413T5 3h4q.825 0 1.413.588T11 5v4q0 .825-.588 1.413T9 11H5Zm0 10q-.825 0-1.413-.588T3 19v-4q0-.825.588-1.413T5 13h4q.825 0 1.413.588T11 15v4q0 .825-.588 1.413T9 21H5Zm10-10q-.825 0-1.413-.588T13 9V5q0-.825.588-1.413T15 3h4q.825 0 1.413.588T21 5v4q0 .825-.588 1.413T19 11h-4Zm0 10q-.825 0-1.413-.588T13 19v-4q0-.825.588-1.413T15 13h4q.825 0 1.413.588T21 15v4q0 .825-.588 1.413T19 21h-4Z"/>`),
		g.Group(children),
	)
}