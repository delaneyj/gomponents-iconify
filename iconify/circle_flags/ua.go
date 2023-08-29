package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ua(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsUa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUa0)"><path fill="#ffda44" d="m0 256l258-39.4L512 256v256H0z"/><path fill="#338af3" d="M0 0h512v256H0z"/></g>`),
		g.Group(children),
	)
}