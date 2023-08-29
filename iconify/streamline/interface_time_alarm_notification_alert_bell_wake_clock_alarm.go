package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTimeAlarmNotificationAlertBellWakeClockAlarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="8" r="5.5"/><path d="M5.5.5h3M7 .5v2M5.5 6L7 8h2.5M12 2l1 1M2 2L1 3"/></g>`),
		g.Group(children),
	)
}