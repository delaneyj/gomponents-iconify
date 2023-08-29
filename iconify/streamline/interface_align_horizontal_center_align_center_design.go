package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlignHorizontalCenterAlignCenterDesign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="7" height="4" x="3.5" y=".5" rx="1"/><rect width="11" height="4" x="1.5" y="9.5" rx="1"/><path d="M7 4.5v5"/></g>`),
		g.Group(children),
	)
}