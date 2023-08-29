package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingListBulletsPointsBulletUnorderedListListsBullets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="1" cy="2.5" r=".5"/><path d="M4.5 2.5h9"/><circle cx="1" cy="7" r=".5"/><path d="M4.5 7h9"/><circle cx="1" cy="11.5" r=".5"/><path d="M4.5 11.5h9"/></g>`),
		g.Group(children),
	)
}