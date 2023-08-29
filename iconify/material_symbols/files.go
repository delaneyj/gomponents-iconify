package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Files(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.825 13H11V7.825ZM4 20q-.825 0-1.412-.587Q2 18.825 2 18v-3h9q.825 0 1.413-.588Q13 13.825 13 13V4h7q.825 0 1.413.588Q22 5.175 22 6v12q0 .825-.587 1.413Q20.825 20 20 20Zm-2-6V8q0-.825.588-1.412Q3.175 6 4 6h6Z"/>`),
		g.Group(children),
	)
}