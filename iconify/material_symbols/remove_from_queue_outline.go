package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveFromQueueOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 12h8v-2H8v2Zm0 9v-2H4q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v12q0 .825-.588 1.413T20 19h-4v2H8Zm-4-4h16V5H4v12Zm0 0V5v12Z"/>`),
		g.Group(children),
	)
}