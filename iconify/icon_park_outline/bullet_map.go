package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 6v36h36"/><path d="M16 40V9c0-1.105 1.053-2 2.353-2h15.294C34.947 7 36 7.895 36 9v31m-19-7.892h17M22 19h8m-4-5v18.108"/></g>`),
		g.Group(children),
	)
}