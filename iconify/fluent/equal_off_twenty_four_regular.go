package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualOffTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06L6.94 8H3.75a.75.75 0 0 0 0 1.5h4.69l5 5H3.75a.75.75 0 0 0 0 1.5h11.19l5.78 5.78a.75.75 0 0 0 1.06-1.06L3.28 2.22ZM20.25 14.5h-2.568l1.5 1.5h1.068a.75.75 0 0 0 0-1.5ZM11.182 8l1.5 1.5h7.568a.75.75 0 0 0 0-1.5h-9.068Z"/>`),
		g.Group(children),
	)
}