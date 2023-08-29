package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pushbullet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 8.409h5.495a1 1 0 0 1 1 1v29.273a1 1 0 0 1-1 1H6.5a2 2 0 0 1-2-2V10.408a2 2 0 0 1 2-2Zm12.636-.091h11.743c16.828.033 16.828 31.337 0 31.273H19.136a1 1 0 0 1-1-1V9.318a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}