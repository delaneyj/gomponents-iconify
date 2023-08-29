package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Je(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsJe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsJe0)"><path fill="#eee" d="M0 47.1L47 0h417.8L512 47.2v417.7L464.9 512H47L0 464.9z"/><path fill="#d80027" d="M0 0v47.1L208.8 256L0 464.9V512h47.1L256 303.2L464.9 512H512v-47L303.1 256L512 47.2V0h-47.2L256 208.9L47 0z"/><path fill="#ffda44" d="M211.5 78L256 89l44.5-11V40l-17.8 9L256 22.3L229.3 49l-17.8-9z"/><path fill="#d80027" d="M211.5 78v27.7c0 34.1 44.5 44.6 44.5 44.6s44.5-10.5 44.5-44.6V78z"/></g>`),
		g.Group(children),
	)
}