package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17V7h2v4h4V7h2v10H9v-4H5v4H3Zm12 0q-.825 0-1.413-.588T13 15V9q0-.825.588-1.413T15 7h6v2h-6v2h4q.825 0 1.413.588T21 13v2q0 .825-.588 1.413T19 17h-4Zm0-4v2h4v-2h-4Z"/>`),
		g.Group(children),
	)
}