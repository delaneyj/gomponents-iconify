package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentMultipleThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 3a4.001 4.001 0 0 0-3.727 2.544A6.06 6.06 0 0 1 7.5 5.5h13a6 6 0 0 1 6 6v9.21A4.001 4.001 0 0 0 29 17V9.5A6.5 6.5 0 0 0 22.5 3h-12ZM3 11.5A4.5 4.5 0 0 1 7.5 7h13a4.5 4.5 0 0 1 4.5 4.5v10a4.5 4.5 0 0 1-4.5 4.5h-5.64l-4.403 3.65C9.48 30.462 8 29.767 8 28.497V26H7a4 4 0 0 1-4-4V11.5Z"/>`),
		g.Group(children),
	)
}