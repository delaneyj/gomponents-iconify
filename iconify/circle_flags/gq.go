package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsGq0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGq0)"><path fill="#eee" d="M41.3 121.9L512 167v178L43.8 391.3z"/><path fill="#6da544" d="M0 0h512v167H111z"/><path fill="#d80027" d="M111 345h401v167H0z"/><path fill="#0052b4" d="M0 0v512l256-256z"/><path fill="#ff9811" d="M334 257.1h22.2v32.3h-22.3z"/><path fill="#6da544" d="M367.3 245a22.3 22.3 0 1 0-44.5 0a11.1 11.1 0 1 0 0 22.1h44.5a11.1 11.1 0 1 0 0-22.2z"/></g>`),
		g.Group(children),
	)
}