package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuTa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsRuTa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuTa0)"><path fill="#eee" d="m0 224l256-32l256 32v64l-256 32L0 288Z"/><path fill="#496e2d" d="M0 0h512v224H0z"/><path fill="#d80027" d="M0 288h512v224H0z"/></g>`),
		g.Group(children),
	)
}