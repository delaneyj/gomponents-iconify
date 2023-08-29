package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingSliderHorizontalSquareAdjustControlsFaderHorizontalSettingsSliderSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><path d="M3 4.5h5"/><circle cx="9.5" cy="4.5" r="1.5"/><path d="M11 9.5H8m-3 0H3"/><circle cx="6.5" cy="9.5" r="1.5"/></g>`),
		g.Group(children),
	)
}