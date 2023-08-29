package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CnXj(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsCnXj0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCnXj0)"><path fill="#338af3" d="M0 0h512v512H0Z"/><path fill="#eee" d="m312 256l116-38l-72 99V195l72 99zm8 69a128 128 0 1 1 0-137a102 102 0 1 0 0 137z"/></g>`),
		g.Group(children),
	)
}