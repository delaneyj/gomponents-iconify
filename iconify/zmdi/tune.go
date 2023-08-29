package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tune(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 299h128v42H0v-42zM0 43h213v42H0V43zm213 341h-42V256h42v43h171v42H213v43zM85 128h43v128H85v-43H0v-42h85v-43zm299 85H171v-42h213v42zm-128-85V0h43v43h85v42h-85v43h-43z"/>`),
		g.Group(children),
	)
}