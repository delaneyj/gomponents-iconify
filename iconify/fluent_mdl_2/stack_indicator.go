package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackIndicator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1707 853h341v1024H683v-341H341v-341H0V171h1365v341h342v341zM171 341v683h1024V341H171zm341 854v170h1024V683h-171v512H512zm1365 512v-683h-170v512H853v171h1024z"/>`),
		g.Group(children),
	)
}