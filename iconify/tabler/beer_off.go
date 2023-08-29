package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7v1.111c0 1.242.29 2.467.845 3.578l.31.622A8 8 0 0 1 9 15.889V20h6v-4.111a8 8 0 0 1 .045-.85m.953-3.035l.157-.315A8 8 0 0 0 17 8.111V4H8M7 8h1m4 0h5M3 3l18 18"/>`),
		g.Group(children),
	)
}