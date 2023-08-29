package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabUnselected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18h2v2Zm-2-4v-2h2v2H2Zm0-4v-2h2v2H2Zm0-4V6q0-.825.588-1.413T4 4h2v2H4v2H2Zm4 12v-2h2v2H6ZM8 6V4h2v2H8Zm2 14v-2h2v2h-2Zm2-10V4h8q.825 0 1.413.588T22 6v4H12Zm2 10v-2h2v2h-2Zm4 0v-2h2v-2h2v2q0 .825-.588 1.413T20 20h-2Zm2-6v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}