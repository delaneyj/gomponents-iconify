package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.487 14.894L12 3L4.03 20.275c-.07.2-.017.424.135.572c.15.148.374.193.57.116l6.275-2.127M15 19l2 2l4-4"/>`),
		g.Group(children),
	)
}