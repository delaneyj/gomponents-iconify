package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingGaugeDashboardOneBarSpeedTestLoadingDashboardInternetGaugeProgress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="m10 4.5l-3 5m-6 0h12"/></g>`),
		g.Group(children),
	)
}