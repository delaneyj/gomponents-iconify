package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingMenuOneAlternateButtonParallelHorizontalLinesMenuNavigationTwoThick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="11" height="4" x="1.5" y="9.61" rx="1"/><rect width="11" height="4" x="1.5" y=".61" rx="1"/></g>`),
		g.Group(children),
	)
}