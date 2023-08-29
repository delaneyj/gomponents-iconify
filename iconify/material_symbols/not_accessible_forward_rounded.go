package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotAccessibleForwardRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.075 21.9l-1.5-1.475h.575q0 .25-.163.413t-.412.162q-.25 0-.413-.163T17 20.425v-.575L13.15 16H11q-1.1 0-1.7-.938T9.15 13.1l.35-.75l-7.425-7.425q-.3-.3-.3-.712t.3-.713q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3ZM19 16.15l-2.65-2.65H17q.825 0 1.413.588T19 15.5v.65Zm-4.175-4.175L9.85 7h4.1q1.125 0 1.713.913T15.8 9.85l-.975 2.125ZM8 22q-2.075 0-3.538-1.463T3 17q0-2.075 1.463-3.538T8 12v2q-1.25 0-2.125.875T5 17q0 1.25.875 2.125T8 20q1.25 0 2.125-.875T11 17h2q0 2.075-1.463 3.538T8 22Zm8-15.5q-.825 0-1.413-.588T14 4.5q0-.825.588-1.413T16 2.5q.825 0 1.413.588T18 4.5q0 .825-.588 1.413T16 6.5Z"/>`),
		g.Group(children),
	)
}