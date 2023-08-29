package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularAltOneBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 20q-.625 0-1.063-.438T5 18.5v-3q0-.625.438-1.063T6.5 14q.625 0 1.063.438T8 15.5v3q0 .625-.438 1.063T6.5 20Z"/>`),
		g.Group(children),
	)
}