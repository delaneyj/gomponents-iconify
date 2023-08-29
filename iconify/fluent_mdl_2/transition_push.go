package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransitionPush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384h2048v1152H0V384zm128 1024h1792V512H128v896zM2048 128v128H0V128h2048zM0 1664h2048v128H0v-128zm1371-768h421v128h-421l162 163l-90 90l-317-317l317-317l90 90l-162 163zM605 643l317 317l-317 317l-90-90l162-163H256V896h421L515 733l90-90z"/>`),
		g.Group(children),
	)
}