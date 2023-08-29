package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsPw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsPw0)"><path fill="#338af3" d="M0 0h512v512H0z"/><circle cx="200.3" cy="256" r="111.3" fill="#ffda44"/></g>`),
		g.Group(children),
	)
}