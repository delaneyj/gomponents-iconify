package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiProtectedSetupRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.4 16.625l.1-.75q.05-.375.05-.75q0-2-.875-3.663T13.3 8.7l-1.45 1.45q-.25.25-.55.125T11 9.8V3.5q0-.2.15-.35T11.5 3h6.3q.35 0 .475.3t-.125.55L16.7 5.3q1.3 1.175 2.075 2.8t.775 3.525q0 1.575-.525 2.95t-1.45 2.5q-.425.5-.85.325t-.325-.775ZM6.2 21q-.35 0-.475-.3t.125-.55L7.3 18.7q-1.325-1.175-2.087-2.8t-.763-3.525q0-1.575.525-2.95t1.475-2.5q.425-.5.85-.325t.3.775q-.075.35-.112.725t-.038.775q0 2 .888 3.663T10.7 15.3l1.45-1.45q.25-.25.55-.125t.3.475v6.3q0 .2-.15.35t-.35.15H6.2Z"/>`),
		g.Group(children),
	)
}