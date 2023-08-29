package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerKeyboardReturnOneKeyboardArrowReturnEnter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 10.5h12a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1h-4"/><path d="m3.5 7.5l-3 3l3 3"/></g>`),
		g.Group(children),
	)
}