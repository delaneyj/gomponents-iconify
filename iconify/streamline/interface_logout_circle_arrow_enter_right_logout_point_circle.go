package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLogoutCircleArrowEnterRightLogoutPointCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.5 7h7m-2-2l2 2l-2 2"/><path d="M11.7 11.49a6.5 6.5 0 1 1 0-9"/></g>`),
		g.Group(children),
	)
}