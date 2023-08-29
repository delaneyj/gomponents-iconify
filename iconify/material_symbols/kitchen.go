package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kitchen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 8h2V5H8v3Zm0 9h2v-5H8v5Zm-2 5q-.825 0-1.413-.588T4 20v-9h16v9q0 .825-.588 1.413T18 22H6ZM4 9V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v5H4Z"/>`),
		g.Group(children),
	)
}