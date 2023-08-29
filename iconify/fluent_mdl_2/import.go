package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Import(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M557 589L249 898h1543v128H252l305 305l-90 90L6 960l461-461l90 90zm1363-77h128v896h-128V512z"/>`),
		g.Group(children),
	)
}