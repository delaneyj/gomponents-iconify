package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StylusLaserPointer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 22q-1.25 0-2.125-.875T6 19q0-1.25.875-2.125T9 16q1.25 0 2.125.875T12 19q0 1.25-.875 2.125T9 22Zm4.475-3.475q-.15-1.375-1.025-2.425t-2.175-1.425L12.925 12H5.9q-.8 0-1.35-.55T4 10.1q0-.5.263-.938t.687-.712L17.1 1.175q.45-.275.95-.138t.775.588q.275.45.138.938t-.563.762L9 9h9.1q.8 0 1.35.55T20 10.9q0 .45-.112.888t-.438.762l-5.975 5.975Z"/>`),
		g.Group(children),
	)
}