package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 12h16V8H4v4Zm15 10v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v6h-3q-2.075 0-3.538 1.463T14 17v3H4Z"/>`),
		g.Group(children),
	)
}