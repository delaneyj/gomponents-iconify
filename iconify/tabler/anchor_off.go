package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12v9m-8-8a8 8 0 0 0 14.138 5.13m1.44-2.56A7.99 7.99 0 0 0 20 13m1 0h-2M5 13H3m9.866-4.127a3 3 0 1 0-3.737-3.747M3 3l18 18"/>`),
		g.Group(children),
	)
}