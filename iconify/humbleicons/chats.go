package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9.882 15C13.261 15 16 12.538 16 9.5S13.261 4 9.882 4C6.504 4 3.765 6.462 3.765 9.5c0 .818.198 1.594.554 2.292L3 15l3.824-.736A6.62 6.62 0 0 0 9.882 15Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.804 18.124a6.593 6.593 0 0 0 3.314.876a6.623 6.623 0 0 0 3.059-.736L21 19l-1.32-3.208a5.02 5.02 0 0 0 .555-2.292c0-1.245-.46-2.393-1.235-3.315c-.749-.89-1.792-1.569-3-1.92"/><circle r="1" fill="currentColor" transform="matrix(-1 0 0 1 13 9.5)"/><circle r="1" fill="currentColor" transform="matrix(-1 0 0 1 10 9.5)"/><circle r="1" fill="currentColor" transform="matrix(-1 0 0 1 7 9.5)"/></g>`),
		g.Group(children),
	)
}