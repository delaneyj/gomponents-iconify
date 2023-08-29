package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransitionPop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384h2048v1152H0V384zm128 1024h1792V512H128v896zm1389-765l318 317l-318 317l-90-90l162-163H475l162 163l-90 90l-317-317l317-317l90 90l-162 163h1114l-162-163l90-90zm531-515v128H0V128h2048zM0 1664h2048v128H0v-128z"/>`),
		g.Group(children),
	)
}