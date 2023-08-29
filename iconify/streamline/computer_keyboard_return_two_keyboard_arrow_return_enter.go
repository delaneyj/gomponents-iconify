package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerKeyboardReturnTwoKeyboardArrowReturnEnter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 9.5h9a4 4 0 0 0 0-8h-3"/><path d="m3.5 6.5l-3 3l3 3"/></g>`),
		g.Group(children),
	)
}