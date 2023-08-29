package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceContentBookPagePagesContentBooksBookOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8 12.5a1 1 0 0 1-1.45.89l-4-2A1 1 0 0 1 2 10.5V.5l5.45 2.72a1 1 0 0 1 .55.9Z"/><path d="M2 .5h9a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H8"/></g>`),
		g.Group(children),
	)
}