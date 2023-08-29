package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Teapod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.457 14.136h10.617s10.07 13.174-4.464 20.58h-6.153"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.038 17.577c5.712-12.724 13.404 6.418.114 11.35M26.248 14.136H15.63S5.56 27.31 20.095 34.716h6.153"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.687 20.36a5.78 5.78 0 0 1-4.631-4.814C7.585 13.79 6.506 14 6.506 14H3.5s1.335.113 2.514 6.67a9.967 9.967 0 0 0 7.334 7.873"/>`),
		g.Group(children),
	)
}