package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermDataSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 20q.825 0 1.413-.588T21 18q0-.825-.588-1.413T19 16q-.825 0-1.413.588T17 18q0 .825.588 1.413T19 20Zm-1 3l-.3-1.5q-.3-.125-.563-.263T16.6 20.9l-1.45.45l-1-1.7l1.15-1q-.05-.3-.05-.65t.05-.65l-1.15-1l1-1.7l1.45.45q.275-.2.537-.337t.563-.263L18 13h2l.3 1.5q.3.125.563.263t.537.337l1.45-.45l1 1.7l-1.15 1q.05.3.05.65t-.05.65l1.15 1l-1 1.7l-1.45-.45q-.275.2-.537.338t-.563.262L20 23h-2ZM2 22L21.975 2.025v9.525q-.65-.325-1.463-.488t-1.537-.162q-2.95 0-5.025 2.075T11.875 18q0 1.125.3 2.1T13.1 22H2Z"/>`),
		g.Group(children),
	)
}