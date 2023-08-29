package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlertRadioActiveTwoWarningRadioactiveRadiationEmergencyDangerSafety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M8.92 1.25a3.2 3.2 0 1 1-3.84 0"/><path d="M.5 9.17a3.2 3.2 0 1 1 1.92 3.32"/><path d="M11.58 12.49a3.19 3.19 0 1 1 1.92-3.32"/></g>`),
		g.Group(children),
	)
}