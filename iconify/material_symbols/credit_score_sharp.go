package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditScoreSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.95 22l-4.25-4.25l1.4-1.4l2.85 2.8l5.65-5.65l1.4 1.45L14.95 22ZM4 12h16V8H4v4Zm-2 8V4h20v8h-2.725l-4.325 4.325l-2.825-2.825l-4.25 4.25l.225.25v2H2Z"/>`),
		g.Group(children),
	)
}