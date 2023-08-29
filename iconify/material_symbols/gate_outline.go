package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 17V7h2v10H2Zm3 2V9q0-1.65 1.175-2.825T9 5h6q1.65 0 2.825 1.175T19 9v10H5Zm15-2V7h2v10h-2Zm-8-5Zm-5 5h4v-4H9v-2h2V7H9q-.825 0-1.413.588T7 9v8Zm6 0h4V9q0-.825-.588-1.413T15 7h-2v4h2v2h-2v4Z"/>`),
		g.Group(children),
	)
}