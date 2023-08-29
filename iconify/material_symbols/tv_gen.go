package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvGen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-2q-.825 0-1.413-.588T2 17V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v11q0 .825-.588 1.413T20 19v2h-1l-.65-2H5.675L5 21H4Z"/>`),
		g.Group(children),
	)
}