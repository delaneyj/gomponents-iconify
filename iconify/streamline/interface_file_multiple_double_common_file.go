package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceFileMultipleDoubleCommonFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="11" x="2" y="2.5" rx="1"/><path d="M4 5h4M4 7.5h4M4 10h2M4.5.5H11a1 1 0 0 1 1 1V11"/></g>`),
		g.Group(children),
	)
}