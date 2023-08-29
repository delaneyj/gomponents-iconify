package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1013 588l-456 419q-18 18-43.5 18t-43.5-18L13 588Q2 577 1 558.5t7-32T28 513h229v-64h128v-64H257V257h128v64h128v-64H385v-65H257V64h128v64h64V0h128v128h-64v129h128V128h64V0h129v128h-65v129H641v64h129v192h228q11 0 19.5 13.5t7.5 32t-12 29.5zM641 385H513v64h128v-64z"/>`),
		g.Group(children),
	)
}