package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsGf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsGf0)"><path fill="#6da544" d="m0 0l216.9 301.6L512 512V0z"/><path fill="#ffda44" d="m0 0l512 512H0z"/><path fill="#d80027" d="m256 121l90 270l-234-168h288L166 391z"/></g>`),
		g.Group(children),
	)
}