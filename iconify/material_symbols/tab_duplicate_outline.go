package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDuplicateOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20h2v2Zm-2-4v-2h2v2H2Zm0-4v-2h2v2H2Zm0-4V8h2v2H2Zm0-4q0-.825.588-1.413T4 4v2H2Zm4 16v-2h2v2H6Zm4 0v-2h2v2h-2Zm4 0v-2h2v2h-2Zm4 0v-2h2q0 .825-.588 1.413T18 22ZM8 18q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm0-2h12V8h-7V4H8v12ZM8 4v12V4Z"/>`),
		g.Group(children),
	)
}