package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsApplicationsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17h2l.3-1.5q.3-.125.563-.263t.537-.337l1.45.45l1-1.7l-1.15-1q.05-.35.05-.65t-.05-.65l1.15-1l-1-1.7l-1.45.45q-.275-.2-.537-.338T13.3 8.5L13 7h-2l-.3 1.5q-.3.125-.563.263T9.6 9.1l-1.45-.45l-1 1.7l1.15 1q-.05.35-.05.65t.05.65l-1.15 1l1 1.7l1.45-.45q.275.2.538.338t.562.262L11 17Zm1-3q-.825 0-1.413-.588T10 12q0-.825.588-1.413T12 10q.825 0 1.413.588T14 12q0 .825-.588 1.413T12 14Zm-7 7q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}