package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditBombDeleteBombRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="6.5" cy="7.5" r="6"/><path d="m13.5.5l-2.76 2.76M3.5 7.5a3 3 0 0 1 3-3"/></g>`),
		g.Group(children),
	)
}