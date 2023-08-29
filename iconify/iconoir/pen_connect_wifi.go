package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenConnectWifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 9.76l.01-.011M3 6.25c2.5-3 7.5-3 10 0m-8 2c1.5-2 4.5-2 6 0m6.5-1.75l1-1a2.121 2.121 0 0 1 3 0v0a2.121 2.121 0 0 1 0 3l-1 1m-3-3L6.696 17.304a1 1 0 0 0-.263.465L5.5 21.5l3.731-.933a1 1 0 0 0 .465-.263L20.5 9.5m-3-3l3 3"/>`),
		g.Group(children),
	)
}