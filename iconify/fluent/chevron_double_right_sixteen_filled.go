package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.7 4.26a.75.75 0 1 1 1.1-1.02l4 4.25a.75.75 0 0 1 0 1.02l-4 4.25a.75.75 0 1 1-1.1-1.02L11.226 8L7.7 4.26Zm-4 0a.75.75 0 1 1 1.1-1.02l4 4.25a.75.75 0 0 1 0 1.02l-4 4.25a.75.75 0 1 1-1.1-1.02L7.227 8L3.7 4.26Z"/>`),
		g.Group(children),
	)
}