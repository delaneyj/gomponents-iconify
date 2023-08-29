package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ewe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsEwe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEwe0)"><path fill="#d80027" d="m0 167l256-32l256 32v178l-256 32L0 345Z"/><path fill="#6da544" d="M0 0h512v167H0zm0 345h512v167H0z"/><path fill="#ffda44" d="m110 200l36 112l-95-69h117l-94 69zm146 0l36 112l-94-69h117l-95 69zm146 0l36 112l-94-69h117l-95 69z"/></g>`),
		g.Group(children),
	)
}