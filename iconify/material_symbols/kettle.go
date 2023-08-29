package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kettle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19V6L3 2h15v3h2q.825 0 1.413.588T22 7v5q0 .825-.588 1.413T20 14h-2v5H6Zm12-7h2V7h-2v5Zm-6 4h3V5h-3v11Zm-9 6v-2h18v2H3Z"/>`),
		g.Group(children),
	)
}