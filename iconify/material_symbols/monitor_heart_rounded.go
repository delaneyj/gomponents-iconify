package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorHeartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18v-5h5.375L9.1 16.45q.125.275.375.413T10 17q.275 0 .525-.138t.375-.412l3.1-6.2l1.1 2.2q.125.275.375.413T16 13h6v5q0 .825-.588 1.413T20 20H4Zm6-6.25l-1.1-2.2q-.125-.275-.375-.413T8 11H2V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v5h-5.375L14.9 7.55q-.125-.275-.375-.388T14 7.05q-.275 0-.525.113t-.375.387l-3.1 6.2Z"/>`),
		g.Group(children),
	)
}