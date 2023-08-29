package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceSettingMenuHorizontalSquareNavigationDotsThreeSquareButtonHorizontalMenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="13" height="13" x=".5" y=".5" rx="1"/><circle cx="7" cy="7" r=".5"/><circle cx="4" cy="7" r=".5"/><circle cx="10" cy="7" r=".5"/></g>`),
		g.Group(children),
	)
}