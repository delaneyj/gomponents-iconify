package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19 4H35L31 36L32.2361 36.618C33.93 37.465 35 39.1963 35 41.0902V44H10V42L23 36L19 4Z"/><path d="M20 12H34"/></g>`),
		g.Group(children),
	)
}