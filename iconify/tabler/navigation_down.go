package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.528 12.815L12 3L4.03 20.275c-.07.2-.017.424.135.572c.15.148.374.193.57.116L12 18.5m7-2.5v6m3-3l-3 3l-3-3"/>`),
		g.Group(children),
	)
}