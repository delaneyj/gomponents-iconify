package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dresser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V5q0-.825.588-1.413T6 3h12q.825 0 1.413.588T20 5v16h-2v-2H6v2H4Zm2-10h5V5H6v6Zm7-4h5V5h-5v2Zm0 4h5V9h-5v2Zm-3 5h4v-2h-4v2Z"/>`),
		g.Group(children),
	)
}