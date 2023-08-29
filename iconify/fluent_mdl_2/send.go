package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Send(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1997 960L18 1843l220-883L18 77l1979 883zM206 301l149 598h1190L206 301zm147 726l-147 592l1327-592H353z"/>`),
		g.Group(children),
	)
}