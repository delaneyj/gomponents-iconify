package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoAwesomeMotion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.825 0-1.413-.588T10 20v-8q0-.825.588-1.413T12 10h8q.825 0 1.413.588T22 12v8q0 .825-.588 1.413T20 22h-8Zm-6-4V8q0-.825.588-1.413T8 6h10v2H8v10H6Zm-4-4V4q0-.825.588-1.413T4 2h10v2H4v10H2Z"/>`),
		g.Group(children),
	)
}