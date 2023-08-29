package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InTg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsInTg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsInTg0)"><path fill="#eee" d="M0 0h512v128l-256 64L0 128Z"/><path fill="#338af3" d="M0 128h512v128l-256 64L0 256Z"/><path fill="#ff9811" d="M0 256h512v128l-256 64L0 384Z"/><path fill="#496e2d" d="M0 384h512v128H0z"/></g>`),
		g.Group(children),
	)
}