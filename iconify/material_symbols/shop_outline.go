package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.5 18l7-4.5l-7-4.5v9ZM4 21q-.825 0-1.413-.588T2 19V6h6V4q0-.825.588-1.413T10 2h4q.825 0 1.413.588T16 4v2h6v13q0 .825-.588 1.413T20 21H4Zm0-2h16V8H4v11Zm6-13h4V4h-4v2ZM4 19V8v11Z"/>`),
		g.Group(children),
	)
}