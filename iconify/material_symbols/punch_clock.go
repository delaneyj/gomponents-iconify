package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PunchClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20V8q0-.825.588-1.413T5 6h1V1h12v5h1q.825 0 1.413.588T21 8v12q0 .825-.588 1.413T19 22H5ZM8 6h8V3H8v3Zm4 8Zm0 4.2q1.725 0 2.963-1.238T16.2 14q0-1.725-1.238-2.963T12 9.8q-1.725 0-2.963 1.238T7.8 14q0 1.725 1.238 2.963T12 18.2Zm1.15-2.35L11.5 14.2v-2.7h1v2.3l1.35 1.35l-.7.7Z"/>`),
		g.Group(children),
	)
}