package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlertRadioActiveThreeWarningRadioactiveRadiationEmergencyDangerSafety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7.49" r="3.5"/><path d="M9.35 1.31a3.32 3.32 0 1 1-4.7 0"/><path d="M.5 8.5a3.32 3.32 0 1 1 2.35 4.07"/><path d="M11.15 12.57A3.32 3.32 0 1 1 13.5 8.5"/></g>`),
		g.Group(children),
	)
}