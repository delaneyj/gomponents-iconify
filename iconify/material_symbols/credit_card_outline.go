package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 6v12q0 .825-.588 1.413T20 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6ZM4 8h16V6H4v2Zm0 4v6h16v-6H4Zm0 6V6v12Z"/>`),
		g.Group(children),
	)
}