package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsNe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsNe0)"><path fill="#eee" d="m0 144.7l255.3-36.5L512 144.7v222.6L250.5 407L0 367.3z"/><path fill="#ff9811" d="M0 0h512v144.7H0z"/><path fill="#6da544" d="M0 367.3h512V512H0z"/><circle cx="256" cy="256.1" r="89" fill="#ff9811"/></g>`),
		g.Group(children),
	)
}