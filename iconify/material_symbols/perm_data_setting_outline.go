package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermDataSettingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22L21.975 2.025v9.525q-.475-.225-.975-.363t-1.025-.212V6.85L6.825 20h5.325q.15.55.4 1.05t.55.95H2Zm4.825-2l13.15-13.15l-6.575 6.575L6.825 20ZM18 23l-.3-1.5q-.3-.125-.563-.263T16.6 20.9l-1.45.45l-1-1.7l1.15-1q-.05-.3-.05-.65t.05-.65l-1.15-1l1-1.7l1.45.45q.275-.2.537-.337t.563-.263L18 13h2l.3 1.5q.3.125.563.263t.537.337l1.45-.45l1 1.7l-1.15 1q.05.3.05.65t-.05.65l1.15 1l-1 1.7l-1.45-.45q-.275.2-.537.338t-.563.262L20 23h-2Zm1-3q.825 0 1.413-.588T21 18q0-.825-.588-1.413T19 16q-.825 0-1.413.588T17 18q0 .825.588 1.413T19 20Z"/>`),
		g.Group(children),
	)
}