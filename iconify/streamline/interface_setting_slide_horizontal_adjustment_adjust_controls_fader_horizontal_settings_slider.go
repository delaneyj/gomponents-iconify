package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingSlideHorizontalAdjustmentAdjustControlsFaderHorizontalSettingsSlider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="5.5" cy="7" r="2"/><path d="M1.5 7h-1m9 0h4"/></g>`),
		g.Group(children),
	)
}