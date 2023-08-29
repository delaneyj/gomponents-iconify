package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingSliderVerticalAdjustmentAdjustControlsFaderVerticalSettingsSlider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2" cy="4.5" r="1.5"/><path d="M2 6v7.5m0-13V3"/><circle cx="12" cy="4.5" r="1.5"/><path d="M12 3V.5m0 13V6"/><circle cx="7" cy="7" r="1.5"/><path d="M7 .5v5m0 3v5"/></g>`),
		g.Group(children),
	)
}