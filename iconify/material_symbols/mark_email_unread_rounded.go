package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkEmailUnreadRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 8q-1.25 0-2.125-.875T16 5q0-1.25.875-2.125T19 2q1.25 0 2.125.875T22 5q0 1.25-.875 2.125T19 8ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h10.1q-.1.5-.1 1t.1 1q.175.8.575 1.488T15.65 8.7L12 11L5.3 6.8q-.425-.275-.863-.025T4 7.525q0 .225.1.413t.3.312l7.075 4.425q.25.15.525.15t.525-.15l4.75-2.975q.425.15.85.225T19 10q.8 0 1.575-.25T22 9v9q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}