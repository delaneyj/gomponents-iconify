package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 16a3 3 0 0 1 3 3v.715C24 23.292 19.79 26 14 26S4 23.433 4 19.715V19a3 3 0 0 1 3-3h14Zm0 1.5H7a1.5 1.5 0 0 0-1.493 1.355L5.5 19v.715c0 2.674 3.389 4.785 8.5 4.785c4.926 0 8.355-2.105 8.495-4.624l.005-.161V19a1.5 1.5 0 0 0-1.355-1.493L21 17.5ZM14 2a6 6 0 1 1 0 12a6 6 0 0 1 0-12Zm0 1.5a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9Z"/>`),
		g.Group(children),
	)
}