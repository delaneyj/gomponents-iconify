package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.1 21.875l-6.375-6.35L7.25 21H3v-4.25l5.475-5.475L2.1 4.9l1.45-1.425l16.975 17l-1.425 1.4Zm-3.525-9.2l-4.25-4.25L13.6 6.15l4.25 4.25l-2.275 2.275Zm3.725-3.75l-4.25-4.2l1.4-1.4q.575-.575 1.413-.575t1.412.575l1.4 1.4q.575.575.6 1.388t-.55 1.387L19.3 8.925Z"/>`),
		g.Group(children),
	)
}