package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrunchDining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 22q-.225 0-.363-.138T2 21.5V20h14v1.5q0 .225-.138.363T15.5 22h-13ZM2 18v-1.5q0-.225.138-.363T2.5 16H7v-2h4v2h4.5q.225 0 .363.138T16 16.5V18H2Zm16 4v-6.1q-.9-1.025-1.45-2.025T16 11.45V2h6v9.45q0 1.425-.537 2.438T20 15.9V20h2v2h-4Zm0-14h2V4h-2v4Z"/>`),
		g.Group(children),
	)
}