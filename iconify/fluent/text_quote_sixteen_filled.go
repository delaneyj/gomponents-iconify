package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextQuoteSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.646 11.146c1.387-1.386 2.008-2.76 2.24-4.352A2 2 0 1 1 13 5c-.001 2.592-.528 4.734-2.647 6.854a.5.5 0 0 1-.708-.708Zm-6 0c1.387-1.386 2.009-2.76 2.24-4.352A2 2 0 1 1 7 5c-.001 2.592-.528 4.734-2.647 6.854a.5.5 0 0 1-.708-.708Z"/>`),
		g.Group(children),
	)
}