package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsGh0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGh0)"><path fill="#ffda44" d="m0 167l256-32l256 32v178l-256 32L0 345Z"/><path fill="#d80027" d="M0 0h512v167H0Z"/><path fill="#496e2d" d="M0 345h512v167H0Z"/><path fill="#333" d="m198 345l151-109H163l151 109l-58-178Z"/></g>`),
		g.Group(children),
	)
}