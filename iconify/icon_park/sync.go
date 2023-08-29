package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M30 7.979L19 8.00006V18.0001H4V32H20.9995"/><path d="M20 42L31 41.9789V31.5789H44V18H30.3448"/></g>`),
		g.Group(children),
	)
}