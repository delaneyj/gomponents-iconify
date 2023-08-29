package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EcW(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsEcW0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEcW0)"><path fill="#eee" d="m0 167l254.6-36.6L512 166.9v178l-254.6 36.4L0 344.9z"/><path fill="#6da544" d="M0 0h512v166.9H0z"/><path fill="#0052b4" d="M0 344.9h512V512H0z"/></g>`),
		g.Group(children),
	)
}