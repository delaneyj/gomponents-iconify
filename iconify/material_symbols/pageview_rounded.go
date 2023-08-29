package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageviewRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm7.5-4q.65 0 1.25-.175t1.15-.525l1.725 1.725q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7L15.3 13.9q.35-.55.525-1.15T16 11.5q0-1.875-1.313-3.188T11.5 7Q9.625 7 8.312 8.313T7 11.5q0 1.875 1.313 3.188T11.5 16Zm0-2q-1.05 0-1.775-.725T9 11.5q0-1.05.725-1.775T11.5 9q1.05 0 1.775.725T14 11.5q0 1.05-.725 1.775T11.5 14Z"/>`),
		g.Group(children),
	)
}