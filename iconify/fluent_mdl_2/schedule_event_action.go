package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScheduleEventAction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 128h384v512H128v1152h896v128H0V128h384V0h128v128h1024V0h128v128zm256 384V256h-256v128h-128V256H512v128H384V256H128v256h1792zm97 256l-238 384h269l-672 896h-264l256-512h-256l387-768h518zm-225 512h-243l238-384h-209l-258 512h256l-241 482l457-610z"/>`),
		g.Group(children),
	)
}