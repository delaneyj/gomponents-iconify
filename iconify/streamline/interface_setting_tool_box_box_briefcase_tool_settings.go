package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingToolBoxBoxBriefcaseToolSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="9" x=".5" y="4.5" rx="1"/><path d="M.5 8.5h13M7 7.5v2m3-5a3 3 0 0 0-3-3h0a3 3 0 0 0-3 3"/></g>`),
		g.Group(children),
	)
}