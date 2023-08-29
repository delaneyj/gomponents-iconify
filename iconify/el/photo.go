package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 30.981v1138.037h1200V30.981H0zm148.096 143.556h903.809v624.097H148.096V174.537z"/>`),
		g.Group(children),
	)
}