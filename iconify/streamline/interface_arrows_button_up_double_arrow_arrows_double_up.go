package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsButtonUpDoubleArrowArrowsDoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 7.54L6.65 1.4a.48.48 0 0 1 .7 0l6.15 6.14"/><path d="M.5 12.75L6.65 6.6a.5.5 0 0 1 .7 0l6.15 6.15"/></g>`),
		g.Group(children),
	)
}