package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AfEmirate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsAfEmirate0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAfEmirate0)"><path fill="#eee" d="M0 0h512v512H0Z"/><path fill="#333" d="M131 189v78c0 12-10 22-22 22v34c30 0 55-25 55-56v-78zm50 0v45a39 39 0 0 0 39 39a39 39 0 0 0 22-7a39 39 0 0 0 22 7l5-1a22 22 0 0 1-21 17v34c30 0 55-25 55-56v-78h-33v45a6 6 0 0 1-11 0v-45h-34v45a6 6 0 0 1-11 0v-45zm139 0v78h33v-78zm50 0v78c0 12-10 22-22 22v34c31 0 56-25 56-56v-78ZM181 289v34h50v-34z"/></g>`),
		g.Group(children),
	)
}