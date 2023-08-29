package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeedbackResponseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1408H704l-448 448v-448H0V128h2048zm-421 429l-136-136l-659 659l-275-275l-136 136l411 411l795-795z"/>`),
		g.Group(children),
	)
}