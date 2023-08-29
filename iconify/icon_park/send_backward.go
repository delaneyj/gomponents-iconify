package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="36" height="8" fill="#2F88FF" transform="matrix(1 0 0 -1 6 14)"/><rect width="36" height="8" fill="#2F88FF" transform="matrix(1 0 0 -1 6 28)"/><path stroke-linecap="round" d="M30 36L24 42L18 36V36"/><path stroke-linecap="round" d="M24 42V28"/><path stroke-linecap="round" d="M24 14V20"/></g>`),
		g.Group(children),
	)
}