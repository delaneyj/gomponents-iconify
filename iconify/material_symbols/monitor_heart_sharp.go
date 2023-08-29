package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorHeartSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-7h5.375L9.1 16.45q.125.275.375.413T10 17q.275 0 .525-.138t.375-.412l3.1-6.2l1.1 2.2q.125.275.375.413T16 13h6v7H2Zm8-6.25l-1.1-2.2q-.125-.275-.375-.413T8 11H2V4h20v7h-5.375L14.9 7.55q-.125-.275-.375-.388T14 7.05q-.275 0-.525.113t-.375.387l-3.1 6.2Z"/>`),
		g.Group(children),
	)
}