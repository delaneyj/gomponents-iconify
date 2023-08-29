package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.929 13.314l2.121 2.121l1.415-1.414l-2.122-2.122l2.122-2.12l2.828 2.828l1.414-1.415L9.88 8.364L12 6.243l2.121 2.12l1.415-1.413l-2.122-2.122L16.243 2a1 1 0 0 1 1.414 0l4.95 4.95a1 1 0 0 1 0 1.414l-14.85 14.85a1 1 0 0 1-1.414 0l-4.95-4.95a1 1 0 0 1 0-1.415l3.536-3.535Z"/>`),
		g.Group(children),
	)
}