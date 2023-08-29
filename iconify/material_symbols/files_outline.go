package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.412-.587Q2 18.825 2 18V8q0-.825.588-1.412Q3.175 6 4 6h6l2-2h8q.825 0 1.413.588Q22 5.175 22 6v12q0 .825-.587 1.413Q20.825 20 20 20Zm1.825-7H11V7.825ZM4 12l4-4H4Zm0 3v3h16V6h-7v7q0 .825-.587 1.412Q11.825 15 11 15Zm7-4Z"/>`),
		g.Group(children),
	)
}