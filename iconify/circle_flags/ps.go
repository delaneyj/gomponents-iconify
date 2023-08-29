package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsPs0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPs0)"><path fill="#eee" d="M41.3 121.9L512 167v178L43.8 391.3z"/><path fill="#333" d="M0 0h512v167H111z"/><path fill="#6da544" d="M111 345h401v167H0z"/><path fill="#d80027" d="M0 0v512l256-256z"/></g>`),
		g.Group(children),
	)
}