package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstitutionChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 .856l10 5.556V9H1V6.412L11 .856ZM4.06 7h13.88L11 3.144L4.06 7ZM6 11v8H4v-8h2Zm6 0v8h-2v-8h2Zm6 0v6h-2v-6h2Zm4.596 6.94l-5.657 5.656l-3.535-3.535l1.414-1.414l2.121 2.12l4.243-4.242l1.414 1.415ZM1 21h11v2H1v-2Z"/>`),
		g.Group(children),
	)
}