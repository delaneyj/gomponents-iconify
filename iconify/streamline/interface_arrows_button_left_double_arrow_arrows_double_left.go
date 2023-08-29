package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsButtonLeftDoubleArrowArrowsDoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7.54.5L1.4 6.65a.48.48 0 0 0 0 .7l6.14 6.15"/><path d="M12.75.5L6.6 6.65a.5.5 0 0 0 0 .7l6.15 6.15"/></g>`),
		g.Group(children),
	)
}