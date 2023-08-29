package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinShake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M42 19L29 6L6 29L19 42L42 19Z"/><path stroke="#fff" d="M16 29L19 32"/><path stroke="#000" d="M30 42L42 30"/><path stroke="#000" d="M6 18L18 6"/></g>`),
		g.Group(children),
	)
}