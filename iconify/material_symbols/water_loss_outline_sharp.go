package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterLossOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20h10l.325-3q-.9 0-1.675-.15t-2.1-.525q-.575-.175-1.2-.25T11.1 16q-1.175 0-2.288.288t-2.137.837L7 20Zm-.55-5q1.125-.5 2.3-.75t2.375-.25q.75 0 1.488.1t1.462.3q1.25.35 1.913.475T17.4 15h.15l1.2-11H5.25l1.2 11ZM5.2 22L3 2h18l-2.2 20H5.2ZM7 20h-.325h10.65H7Z"/>`),
		g.Group(children),
	)
}