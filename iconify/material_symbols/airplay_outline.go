package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 21l6-6l6 6H6Zm-4-4V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v12q0 .825-.588 1.413T20 19h-3v-2h3V5H4v12h3v2H4q-.825 0-1.413-.588T2 17Zm10-5Z"/>`),
		g.Group(children),
	)
}