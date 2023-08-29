package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddCardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v6H4v6h10v2H4ZM4 8h16V6H4v2Zm15 14v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM4 18V6v12Z"/>`),
		g.Group(children),
	)
}