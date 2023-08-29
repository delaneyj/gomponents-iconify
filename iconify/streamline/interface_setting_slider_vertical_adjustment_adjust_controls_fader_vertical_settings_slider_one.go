package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingSliderVerticalAdjustmentAdjustControlsFaderVerticalSettingsSliderOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="5.5" r="2"/><path d="M7 1.5v-1m0 9v4"/></g>`),
		g.Group(children),
	)
}