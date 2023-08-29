package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingSliderVerticalAdjustmentAdjustControlsFaderVerticalSettingsSliderSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1" transform="rotate(-90 7 7)"/><path d="M4.5 11V6"/><circle cx="4.5" cy="4.5" r="1.5"/><path d="M9.5 3v3m0 3v2"/><circle cx="9.5" cy="7.5" r="1.5"/></g>`),
		g.Group(children),
	)
}