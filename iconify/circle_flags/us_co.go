package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsCo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsUsCo0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsCo0)"><path fill="#0052b4" d="M0 0h512v167l-64 89l64 89v167H0V345l64-89l-64-89Z"/><path fill="#eee" d="M0 167h512v178H0z"/><path fill="#d80027" d="M344.3 299.8A128 128 0 0 1 201.8 382A128 128 0 0 1 96 256a128 128 0 0 1 105.8-126a128 128 0 0 1 142.5 82.2L224 256z"/><circle cx="224" cy="256" r="64" fill="#ffda44"/></g>`),
		g.Group(children),
	)
}