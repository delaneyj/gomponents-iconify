package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestStep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1001 1536l-64 128H128v-640h512V640h512V256h768v512h-128V384h-512v384H768v384H256v384h745zm962 327q14 28 14 57q0 26-10 49t-27 41t-41 28t-50 10h-754q-26 0-49-10t-41-27t-28-41t-10-50q0-29 14-57l299-598v-241h-128V896h640v128h-128v241l299 598zm-242-199l-185-369v-271h-128v271l-185 369h498z"/>`),
		g.Group(children),
	)
}