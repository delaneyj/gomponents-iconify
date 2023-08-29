package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FitScreenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 9V6h-3V4h3q.825 0 1.413.588T22 6v3h-2ZM2 9V6q0-.825.588-1.413T4 4h3v2H4v3H2Zm15 11v-2h3v-3h2v3q0 .825-.588 1.413T20 20h-3ZM4 20q-.825 0-1.413-.588T2 18v-3h2v3h3v2H4Zm2-4V8h12v8H6Zm2-2h8v-4H8v4Zm0 0v-4v4Z"/>`),
		g.Group(children),
	)
}