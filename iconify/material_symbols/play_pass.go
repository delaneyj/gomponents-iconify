package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayPass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-.825 0-1.413.588T10 22H7q-.825 0-1.413-.588T5 20V4q0-.825.588-1.413T7 2h3q0 .825.588 1.413T12 4q.825 0 1.413-.588T14 2h3q.825 0 1.413.588T19 4v16q0 .825-.588 1.413T17 22h-3q0-.825-.588-1.413T12 20Z"/>`),
		g.Group(children),
	)
}