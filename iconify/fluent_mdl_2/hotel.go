package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hotel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1152v768h-128v-128H128v128H0v-768h128v113l256-512V128h128v128h1024V128h128v384q27 0 50 10t40 27t28 41t10 50v256q0 39-21 70l149 299v-113h128zm-896-512v256h512V640h-512zM512 384v384h512V640q0-27 10-50t27-40t41-28t50-10h384V384H512zm-248 896h1520l-128-256h-504q-27 0-50-10t-40-27t-28-41t-10-50H455l-191 384zm1656 384v-256H128v256h1792z"/>`),
		g.Group(children),
	)
}