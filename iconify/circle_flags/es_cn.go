package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EsCn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsEsCn0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsEsCn0)"><path fill="#338af3" d="M167 0h178l32.3 257L345 512H167l-25.3-256z"/><path fill="#eee" d="M0 0h166.9v512H0z"/><path fill="#ffda44" d="M344.9 0H512v512H344.9z"/></g>`),
		g.Group(children),
	)
}