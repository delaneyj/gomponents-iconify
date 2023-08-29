package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingSliderHorizontalAdjustmentAdjustControlsFaderHorizontalSettingsSlider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2" cy="2" r="1.5"/><path d="M3.5 2h10"/><circle cx="7" cy="7" r="1.5"/><path d="M.5 7h5m3 0h5"/><circle cx="12" cy="12" r="1.5"/><path d="M10.5 12H.5"/></g>`),
		g.Group(children),
	)
}