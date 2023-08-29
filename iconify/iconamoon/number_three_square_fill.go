package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-14ZM10 6a1 1 0 0 0 0 2h3l-1.8 2.4A1 1 0 0 0 12 12a2 2 0 1 1-1.333 3.491a1 1 0 1 0-1.334 1.49a4 4 0 1 0 4.379-6.597L15.8 7.6A1 1 0 0 0 15 6h-5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}