package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="26" height="40" x="11" y="4" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M22 10h4m-6 28h8"/></g>`),
		g.Group(children),
	)
}