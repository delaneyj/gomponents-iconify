package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentsDisabledOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18q-.825 0-1.413-.588T2 16V4.825L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425l-5.3-5.3H4Zm18 1.125L18.875 16H20V4H6.875l-2-2H20q.825 0 1.413.588T22 4v15.125ZM4 16h9.175l-2-2H6v-2h3.175l-1-1H6V9h.175L4 6.825V16Zm12.875-2l-2-2H18v2h-1.125Zm-3-3l-2-2H18v2h-4.125Zm-3-3l-2-2H18v2h-7.125ZM8.6 11.4Zm4.275-1.4Z"/>`),
		g.Group(children),
	)
}