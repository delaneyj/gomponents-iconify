package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.343 14.728l-2.828 2.828l3.535 3.536L20.485 7.657L16.95 4.12l-2.121 2.122l1.414 1.414l-1.414 1.414l-1.415-1.414l-2.121 2.121l2.121 2.121L12 13.314l-2.12-2.122l-2.122 2.122l1.415 1.414l-1.415 1.414l-1.414-1.414ZM17.657 2l4.95 4.95a1 1 0 0 1 0 1.414l-14.85 14.85a1 1 0 0 1-1.414 0l-4.95-4.95a1 1 0 0 1 0-1.415L16.243 2a1 1 0 0 1 1.414 0Z"/>`),
		g.Group(children),
	)
}