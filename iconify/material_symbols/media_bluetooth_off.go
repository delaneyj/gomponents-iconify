package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaBluetoothOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.8 22.6l-4-4l-.6.6l-.85-.85l.6-.6L11 13.8V17q0 1.65-1.175 2.825T7 21q-1.65 0-2.825-1.175T3 17q0-1.65 1.175-2.825T7 13q.575 0 1.062.138T9 13.55V11.8L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4Zm1.4-4.25l-6.8-6.75l.8-.8l2.75 2.75V9h.6L22 12.4L19.4 15l2.6 2.55l-.8.8Zm-2.05-4.8l1.15-1.15l-1.15-1.1v2.25ZM11 8.15l-2-2V3h6v4h-4v1.15Z"/>`),
		g.Group(children),
	)
}