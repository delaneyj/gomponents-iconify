package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathSymbolsSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 6.25v-1h-1a.75.75 0 0 1 0-1.5h1v-1a.75.75 0 0 1 1.5 0v1h1a.75.75 0 0 1 0 1.5h-1v1a.75.75 0 0 1-1.5 0Zm6-2.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5Zm-7.53 8.97a.75.75 0 1 0 1.06 1.06l1.22-1.22l1.22 1.22a.75.75 0 0 0 1.06-1.06L5.56 11.5l1.22-1.22a.75.75 0 1 0-1.06-1.06L4.5 10.44L3.28 9.22a.75.75 0 0 0-1.06 1.06l1.22 1.22l-1.22 1.22ZM11.5 10a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-2 .25a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5h-4Zm3 2.75a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`),
		g.Group(children),
	)
}