package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoHeightOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M7 42L7 6"/><path stroke-linejoin="round" d="M18 13.9907L23.9954 8L30 14"/><path stroke-linejoin="round" d="M30 34.0093L24.0046 40L18 34"/><path stroke-linejoin="round" d="M24 8V40"/><path d="M41 42L41 6"/></g>`),
		g.Group(children),
	)
}