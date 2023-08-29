package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaBluetoothOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21q-1.65 0-2.825-1.175T3 17q0-1.65 1.175-2.825T7 13q.575 0 1.063.138T9 13.55V3h6v4h-4v10q0 1.65-1.175 2.825T7 21Zm10 0v-4.55l-2.75 2.75l-.85-.85L16.75 15l-3.4-3.35l.85-.85l2.75 2.75V9h.6L21 12.45L18.4 15l2.6 2.55L17.6 21H17Zm1.15-2.3l1.15-1.15l-1.15-1.1v2.25Zm0-5.15l1.15-1.1l-1.15-1.15v2.25Z"/>`),
		g.Group(children),
	)
}