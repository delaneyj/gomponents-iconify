package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsSc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsSc0)"><path fill="#0052b4" d="M0 0v332l150.9-138.5L225.2 0z"/><path fill="#ffda44" d="M273.1 253.3L512 0H225.2L0 332v80.2z"/><path fill="#d80027" d="M512 0L0 412.2v50.4L277.9 390L512 256z"/><path fill="#eee" d="M0 462.6L512 256v133.5l-223.9 78.8L0 488.4z"/><path fill="#6da544" d="m512 389.5l-512 99V512h512z"/></g>`),
		g.Group(children),
	)
}