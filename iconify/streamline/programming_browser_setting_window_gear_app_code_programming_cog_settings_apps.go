package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingBrowserSettingWindowGearAppCodeProgrammingCogSettingsApps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M1.5 13.5a1 1 0 0 1-1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1m-12-10h13M7 6.5V8m-3.03.25l1.3.75m-1.3 2.75l1.3-.75M7 13.5V12m3.03-.25L8.73 11m1.3-2.75L8.73 9"/><circle cx="7" cy="10" r="2"/></g>`),
		g.Group(children),
	)
}