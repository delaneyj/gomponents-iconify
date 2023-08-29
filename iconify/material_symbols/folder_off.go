package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.775 18.95L6.85 4H9.7l2 2H20q.825 0 1.413.588T22 8v10q0 .275-.05.513t-.175.437Zm-1.3 4.35l-3.3-3.3H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4l2 2H3.2L.7 3.5l1.4-1.4l19.8 19.8l-1.425 1.4Z"/>`),
		g.Group(children),
	)
}