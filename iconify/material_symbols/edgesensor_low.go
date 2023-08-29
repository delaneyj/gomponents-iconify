package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgesensorLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 14V7h2v7H2Zm18 3v-7h2v7h-2ZM8 22q-.825 0-1.413-.588T6 20V4q0-.825.588-1.413T8 2h8q.825 0 1.413.588T18 4v16q0 .825-.588 1.413T16 22H8Zm0-5h8V7H8v10Z"/>`),
		g.Group(children),
	)
}