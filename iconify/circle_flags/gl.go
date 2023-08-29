package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsGl0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGl0)"><path fill="#d80027" d="m0 256l259-45.3L512 256v256H0z"/><path fill="#eee" d="M0 0h512v256H0z"/><path fill="#eee" d="M55.7 256a122.4 122.4 0 1 0 244.8 0l-123-24z"/><path fill="#d80027" d="M55.7 256a122.4 122.4 0 1 1 244.8 0z"/></g>`),
		g.Group(children),
	)
}