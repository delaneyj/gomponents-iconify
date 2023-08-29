package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkEmailReadRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v4.35l-6.025 6.025l-1.45-1.45q-.575-.575-1.425-.575t-1.425.575L10.3 16.3q-.575.575-.575 1.425t.575 1.425l.85.85H4Zm8-9L5.3 6.8q-.425-.275-.862-.025T4 7.525q0 .225.1.413t.3.312l7.075 4.425q.25.15.525.15t.525-.15L19.6 8.25q.2-.125.3-.313t.1-.412q0-.5-.437-.75T18.7 6.8L12 11Zm3.95 8.2l4.95-4.95q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-5.65 5.65q-.3.3-.7.3t-.7-.3l-2.85-2.85q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l2.15 2.15Z"/>`),
		g.Group(children),
	)
}