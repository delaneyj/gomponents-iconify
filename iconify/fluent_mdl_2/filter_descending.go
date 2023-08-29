package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterDescending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1795 1667l163-163l90 90l-315 317l-325-325l90-90l169 170v-642h127l1 643zm-643 125v-731l768-768v-37H128v37l768 768v731h256zm896-1445l-768 768v805H768v-805L0 347V128h2048v219z"/>`),
		g.Group(children),
	)
}