package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsMu0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMu0)"><path fill="#6da544" d="m0 378.3l254-37.1l258 37V512H0z"/><path fill="#ffda44" d="m0 256.1l252.2-33.3L512 256v122.4H0z"/><path fill="#0052b4" d="M0 133.7L249.7 97L512 133.7v122.4H0z"/><path fill="#d80027" d="M0 0h512v133.7H0z"/></g>`),
		g.Group(children),
	)
}