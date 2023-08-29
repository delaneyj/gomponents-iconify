package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsLc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsLc0)"><path fill="#338af3" d="M0 0h512v512H0z"/><path fill="#eee" d="M161.4 345h189.2L256 122.4z"/><path fill="#333" d="M194.3 322.8L256 182.4l61.7 140.4z"/><path fill="#ffda44" d="M161.4 345h189.2L256 256z"/></g>`),
		g.Group(children),
	)
}