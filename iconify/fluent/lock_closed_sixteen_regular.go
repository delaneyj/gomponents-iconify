package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClosedSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM5 4h-.5A2.5 2.5 0 0 0 2 6.5v5A2.5 2.5 0 0 0 4.5 14h7a2.5 2.5 0 0 0 2.5-2.5v-5A2.5 2.5 0 0 0 11.5 4H11v-.5a3 3 0 0 0-6 0V4Zm1-.5a2 2 0 1 1 4 0V4H6v-.5ZM11.5 5A1.5 1.5 0 0 1 13 6.5v5a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 11.5v-5A1.5 1.5 0 0 1 4.5 5h7Z"/>`),
		g.Group(children),
	)
}