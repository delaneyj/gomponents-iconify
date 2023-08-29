package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorHeartOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 9V4h20v5h-2V6H4v3H2Zm0 11v-5h2v3h16v-3h2v5H2Zm8-3q.275 0 .525-.138t.375-.412l3.1-6.2l1.1 2.2q.125.275.375.413T16 13h6v-2h-5.375L14.9 7.55q-.275-.5-.9-.5t-.9.5l-3.1 6.2l-1.1-2.2q-.125-.275-.375-.413T8 11H2v2h5.375L9.1 16.45q.125.275.375.413T10 17Zm2-5Z"/>`),
		g.Group(children),
	)
}