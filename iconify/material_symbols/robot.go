package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Robot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12q.825 0 1.413-.588T11 10q0-.825-.588-1.413T9 8q-.825 0-1.413.588T7 10q0 .825.588 1.413T9 12Zm6 0q.825 0 1.413-.588T17 10q0-.825-.588-1.413T15 8q-.825 0-1.413.588T13 10q0 .825.588 1.413T15 12Zm-6 9v-4h2v4H9Zm4 0v-4h2v4h-2Zm-8 0q-.825 0-1.413-.588T3 19V9q0-2.5 1.75-4.25T9 3h6q2.5 0 4.25 1.75T21 9v10q0 .825-.588 1.413T19 21h-2v-4q0-.825-.588-1.413T15 15H9q-.825 0-1.413.588T7 17v4H5Z"/>`),
		g.Group(children),
	)
}