package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsYu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsYu0)"><path fill="#eee" d="m0 167l256-32l256 32v178l-256 32L0 345Z"/><path fill="#d80027" d="M0 345h512v167H0Z"/><path fill="#0052b4" d="M0 0h512v167H0Z"/><path fill="#ffda44" d="m137 413l309-222H66l309 222L256 51Z"/><path fill="#d80027" d="m183 350l189-136H140l188 136l-72-221z"/></g>`),
		g.Group(children),
	)
}