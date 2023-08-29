package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsCm0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCm0)"><path fill="#d80027" d="M144.8 0h222.4l32 260l-32 252H144.8l-32.1-256z"/><path fill="#ffda44" d="m256.1 167l22.1 68h71.5L292 277l22 68l-57.8-42l-57.9 42l22.1-68l-57.8-42H234z"/><path fill="#496e2d" d="M0 0h144.8v512H0z"/><path fill="#ffda44" d="M367.2 0H512v512H367.2z"/></g>`),
		g.Group(children),
	)
}