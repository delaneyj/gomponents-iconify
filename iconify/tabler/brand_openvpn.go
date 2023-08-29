package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandOpenvpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.618 20.243l-2.193-5.602a3 3 0 1 0-2.849 0l-2.193 5.603"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/></g>`),
		g.Group(children),
	)
}