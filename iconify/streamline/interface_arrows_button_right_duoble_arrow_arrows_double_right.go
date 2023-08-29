package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceArrowsButtonRightDuobleArrowArrowsDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m6.46.5l6.14 6.15a.48.48 0 0 1 0 .7L6.46 13.5"/><path d="M1.25.5L7.4 6.65a.5.5 0 0 1 0 .7L1.25 13.5"/></g>`),
		g.Group(children),
	)
}