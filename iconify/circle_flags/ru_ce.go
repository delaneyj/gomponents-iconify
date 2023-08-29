package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuCe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsRuCe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuCe0)"><path fill="#d80027" d="m96 357l208-32l208 32v155H96l-32-78Z"/><path fill="#496e2d" d="M96 0h416v293l-208 32l-208-32l-32-147Z"/><path fill="#eee" d="M0 0v512h96V357h416v-64H96V0Z"/></g>`),
		g.Group(children),
	)
}