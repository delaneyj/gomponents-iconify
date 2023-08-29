package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CribOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14h12v-3h-8V6H8q-.825 0-1.413.588T6 8v6Zm6 6q.5 0 1-.063t1-.187V16h-4v3.75q.5.125 1 .188T12 20Zm0 2q-2 0-3.825-.763t-3.25-2.162L6.35 17.65q.375.375.788.688t.862.587V16H6q-.825 0-1.413-.588T4 14V8q0-1.65 1.175-2.825T8 4h4v5h6q.825 0 1.413.588T20 11v3q0 .825-.588 1.413T18 16h-2v2.925q.45-.275.863-.588t.787-.687l1.425 1.425q-1.425 1.4-3.25 2.163T12 22Zm0-12Z"/>`),
		g.Group(children),
	)
}