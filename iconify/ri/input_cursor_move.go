package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InputCursorMove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2h3V5H8V3h8v2h-3v14h3v2H8ZM18.05 7.05L23 12l-4.95 4.95l-1.414-1.414L20.172 12l-3.536-3.536L18.05 7.05Zm-12.1 0l1.414 1.414L3.828 12l3.536 3.536L5.95 16.95L1 12l4.95-4.95Z"/>`),
		g.Group(children),
	)
}