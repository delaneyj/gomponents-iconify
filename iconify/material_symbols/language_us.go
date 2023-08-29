package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageUs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17q-.825 0-1.413-.588T4 15V7h2v8h3V7h2v8q0 .825-.588 1.413T9 17H6Zm9 0q-.825 0-1.413-.588T13 15v-1h2v1h3v-2h-3q-.825 0-1.413-.588T13 11V9q0-.825.588-1.413T15 7h3q.825 0 1.413.588T20 9v1h-2V9h-3v2h3q.825 0 1.413.588T20 13v2q0 .825-.588 1.413T18 17h-3Z"/>`),
		g.Group(children),
	)
}