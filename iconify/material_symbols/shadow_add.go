package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShadowAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 14v-3h-3V9h3V6h2v3h3v2h-3v3h-2Zm-9 8q-.825 0-1.413-.588T2 20V8q0-.825.588-1.413T4 6h2V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18h-2v2q0 .825-.588 1.413T16 22H4Zm4-6h12V4H8v12Z"/>`),
		g.Group(children),
	)
}