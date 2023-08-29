package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12H4c0-2.209.895-4.208 2.342-5.656m2.382-1.645A8 8 0 0 1 20 12h-4m-4 0v6a2 2 0 1 0 4 0M3 3l18 18"/>`),
		g.Group(children),
	)
}