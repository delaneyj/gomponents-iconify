package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDesktopWarningAlertDesktopDisplayDeviceWarningNotificationComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 2H1a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5v-8A.5.5 0 0 0 13 2h-2.5M6 11l-1 2.5M8 11l1 2.5m-5 0h6m-3-12V5"/><circle cx="7" cy="8" r=".5"/></g>`),
		g.Group(children),
	)
}