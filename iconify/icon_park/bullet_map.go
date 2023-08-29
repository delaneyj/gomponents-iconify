package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 6V42H42"/><path d="M16 40V9C16 7.89543 17.0534 7 18.3529 7H33.6471C34.9466 7 36 7.89543 36 9V40"/><path d="M17 32.1084H34"/><path d="M22 19H30"/><path d="M26 14V32.1082"/></g>`),
		g.Group(children),
	)
}