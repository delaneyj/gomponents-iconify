package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabRecentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 23q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Zm1.65-2.65l.7-.7l-1.85-1.85V15h-1v3.2l2.15 2.15ZM13 10h7v1.3q.55.175 1.05.413t.95.562V6q0-.825-.588-1.413T20 4H4q-.825 0-1.413.588T2 6v12q0 .825.588 1.413T4 20h7.3q-.15-.475-.225-.975T11 18H4V6h9v4Zm-9 8V6v12Z"/>`),
		g.Group(children),
	)
}