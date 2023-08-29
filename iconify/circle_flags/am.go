package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Am(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsAm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsAm0)"><path fill="#0052b4" d="m0 166.9l253-26.7L512 167v178l-261.1 26L0 344.8z"/><path fill="#d80027" d="M0 0h512v166.9H0z"/><path fill="#ff9811" d="M0 344.9h512V512H0z"/></g>`),
		g.Group(children),
	)
}