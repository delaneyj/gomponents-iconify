package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kanuri(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsKanuri0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKanuri0)"><path fill="#ffda44" d="m0 167l256-32l256 32v178l-256 32L0 345Z"/><path fill="#496e2d" d="M0 345h512v167H0Z"/><path fill="#0052b4" d="M0 0h512v167H0Z"/><path fill="#ffda44" d="m373 373l36 112l-94-69h117l-95 69z"/></g>`),
		g.Group(children),
	)
}