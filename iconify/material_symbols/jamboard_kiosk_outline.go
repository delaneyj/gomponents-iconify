package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JamboardKioskOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-2h5v-3H4q-.825 0-1.413-.588T2 14V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v9q0 .825-.588 1.413T20 16h-7v3h5v2H6Zm-2-7h16V5H4v9Zm0 0V5v9Z"/>`),
		g.Group(children),
	)
}