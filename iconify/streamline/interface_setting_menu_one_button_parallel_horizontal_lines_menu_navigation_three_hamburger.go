package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingMenuOneButtonParallelHorizontalLinesMenuNavigationThreeHamburger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 2H.5m13 5H.5m13 5H.5"/>`),
		g.Group(children),
	)
}