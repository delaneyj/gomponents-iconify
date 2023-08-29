package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommuterBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31.992 22c.005-.308.008-.617.008-.927C32 9.515 28.418 4 24 4s-8 5.515-8 17.073c0 .31.003.62.008.927"/><rect width="22" height="22" x="13" y="22" rx="2"/></g>`),
		g.Group(children),
	)
}