package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfacePageControllerSettingsPageSettingSquareTriangleCircleLineCombinationVariation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 1h5v5h-5zm12.75 12.5h-5m0-5h5m-5 2.5h5m.25-5H8L10.75.5L13.5 6z"/><circle cx="3" cy="11" r="2.5"/></g>`),
		g.Group(children),
	)
}