package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgesensorLowOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 14V7h2v7H2Zm18 3v-7h2v7h-2ZM8 22q-.825 0-1.413-.588T6 20V4q0-.825.588-1.413T8 2h8q.825 0 1.413.588T18 4v16q0 .825-.588 1.413T16 22H8Zm8-3H8v1h8v-1ZM8 5h8V4H8v1Zm0 0V4v1Zm0 14v1v-1Zm0-2h8V7H8v10Z"/>`),
		g.Group(children),
	)
}