package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuBa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsRuBa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuBa0)"><path fill="#eee" d="m0 160l256-32l256 32v192l-256 32L0 352Z"/><path fill="#338af3" d="M0 0h512v160H0Z"/><path fill="#496e2d" d="M0 352h512v160H0Z"/><circle cx="256" cy="256" r="64" fill="#ffda44"/></g>`),
		g.Group(children),
	)
}