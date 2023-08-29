package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDesktopSettingSettingsDesktopDisplayDeviceGearCogComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2h1a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H1a.5.5 0 0 1-.5-.5v-8A.5.5 0 0 1 1 2h1m4 9l-1 2.5M8 11l1 2.5m-5 0h6M7.03 2v1.5M4 3.75l1.3.75M4 7.25l1.3-.75M7.03 9V7.5m3.03-.25l-1.3-.75m1.3-2.75l-1.3.75"/><circle cx="7.03" cy="5.5" r="2"/></g>`),
		g.Group(children),
	)
}