package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsSd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSd0)"><path fill="#eee" d="M43.6 109.4L512 144.7v222.6L43.8 397.2z"/><path fill="#d80027" d="M0 0h512v144.7H111z"/><path fill="#333" d="M111 367.3h401V512H0z"/><path fill="#496e2d" d="M0 0v512l256-256z"/></g>`),
		g.Group(children),
	)
}