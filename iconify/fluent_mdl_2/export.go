package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1497 589l308 309H262v128h1540l-305 305l90 90l461-461l-461-461l-90 90zM134 512H6v896h128V512z"/>`),
		g.Group(children),
	)
}