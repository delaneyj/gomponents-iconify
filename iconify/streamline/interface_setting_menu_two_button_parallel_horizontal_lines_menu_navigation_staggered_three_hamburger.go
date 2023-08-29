package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingMenuTwoButtonParallelHorizontalLinesMenuNavigationStaggeredThreeHamburger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 2H6m5 5H3.5m5 5h-8"/>`),
		g.Group(children),
	)
}