package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func De(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsDe0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsDe0)"><path fill="#ffda44" d="m0 345l256.7-25.5L512 345v167H0z"/><path fill="#d80027" d="m0 167l255-23l257 23v178H0z"/><path fill="#333" d="M0 0h512v167H0z"/></g>`),
		g.Group(children),
	)
}