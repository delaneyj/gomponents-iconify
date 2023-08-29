package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingGaugeDashboardTwoBarSpeedTestLoadingDashboardInternetGaugeProgress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10 6.5l-3 5m6 0a6.5 6.5 0 1 0-12 0Z"/>`),
		g.Group(children),
	)
}