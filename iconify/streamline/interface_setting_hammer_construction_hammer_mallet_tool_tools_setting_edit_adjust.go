package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingHammerConstructionHammerMalletToolToolsSettingEditAdjust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="5.66" height="9.66" x="5.67" y=".67" rx="1" transform="rotate(-45 8.498 5.5)"/><path d="M.91 11.09a1.41 1.41 0 0 0 2 2L7.5 8.5l-2-2Z"/></g>`),
		g.Group(children),
	)
}