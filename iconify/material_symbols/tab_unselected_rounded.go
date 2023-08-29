package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabUnselectedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20v-2h2v2h-2ZM8 6V4h2v2H8Zm6 4q-.825 0-1.413-.588T12 8V4h8q.825 0 1.413.588T22 6v4h-8Zm4 10v-2h2v-2h2v2q0 .825-.588 1.413T20 20h-2Zm-8 0v-2h2v2h-2Zm10-6v-2h2v2h-2ZM2 16v-2h2v2H2Zm0-4v-2h2v2H2Zm2 8q-.825 0-1.413-.588T2 18h2v2ZM2 8V6q0-.825.588-1.413T4 4h2v2H4v2H2Zm4 12v-2h2v2H6Z"/>`),
		g.Group(children),
	)
}