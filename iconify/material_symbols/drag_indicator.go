package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragIndicator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20q-.825 0-1.413-.588T7 18q0-.825.588-1.413T9 16q.825 0 1.413.588T11 18q0 .825-.588 1.413T9 20Zm6 0q-.825 0-1.413-.588T13 18q0-.825.588-1.413T15 16q.825 0 1.413.588T17 18q0 .825-.588 1.413T15 20Zm-6-6q-.825 0-1.413-.588T7 12q0-.825.588-1.413T9 10q.825 0 1.413.588T11 12q0 .825-.588 1.413T9 14Zm6 0q-.825 0-1.413-.588T13 12q0-.825.588-1.413T15 10q.825 0 1.413.588T17 12q0 .825-.588 1.413T15 14ZM9 8q-.825 0-1.413-.588T7 6q0-.825.588-1.413T9 4q.825 0 1.413.588T11 6q0 .825-.588 1.413T9 8Zm6 0q-.825 0-1.413-.588T13 6q0-.825.588-1.413T15 4q.825 0 1.413.588T17 6q0 .825-.588 1.413T15 8Z"/>`),
		g.Group(children),
	)
}