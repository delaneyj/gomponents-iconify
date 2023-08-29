package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualizerRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 20q-.825 0-1.413-.588T4 18v-4q0-.825.588-1.413T6 12q.825 0 1.413.588T8 14v4q0 .825-.588 1.413T6 20Zm6 0q-.825 0-1.413-.588T10 18V6q0-.825.588-1.413T12 4q.825 0 1.413.588T14 6v12q0 .825-.588 1.413T12 20Zm6 0q-.825 0-1.413-.588T16 18v-7q0-.825.588-1.413T18 9q.825 0 1.413.588T20 11v7q0 .825-.588 1.413T18 20Z"/>`),
		g.Group(children),
	)
}