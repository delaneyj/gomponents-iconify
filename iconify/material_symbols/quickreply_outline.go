package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickreplyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v6h-2V4H4v13.125L5.15 16H15v2H6l-4 4Zm2-6V4v12Zm15 7v-5h-2v-6h5l-1.7 4h2.2L19 23Z"/>`),
		g.Group(children),
	)
}