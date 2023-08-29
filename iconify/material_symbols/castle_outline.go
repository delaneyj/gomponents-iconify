package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21V9h2v2h2V3h2v2h2V3h2v2h2V3h2v2h2V3h2v8h2V9h2v12h-9v-3q0-.825-.588-1.413T12 16q-.825 0-1.413.588T10 18v3H1Zm2-2h5v-1q0-1.65 1.175-2.825T12 14q1.65 0 2.825 1.175T16 18v1h5v-6h-4V7H7v6H3v6Zm6-7h2V9H9v3Zm4 0h2V9h-2v3Zm-1 1Z"/>`),
		g.Group(children),
	)
}