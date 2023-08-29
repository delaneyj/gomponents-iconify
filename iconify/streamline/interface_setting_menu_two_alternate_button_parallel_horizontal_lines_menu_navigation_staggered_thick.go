package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingMenuTwoAlternateButtonParallelHorizontalLinesMenuNavigationStaggeredThick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="9" height="4" x=".5" y="1.5" rx="1"/><rect width="9" height="4" x="4.5" y="8.5" rx="1"/></g>`),
		g.Group(children),
	)
}