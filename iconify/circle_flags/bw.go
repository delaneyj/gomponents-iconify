package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsBw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBw0)"><path fill="#338af3" d="M0 0h512v178l-31 76.9l31 79.1v178H0V334l37-80.7L0 178z"/><path fill="#333" d="m0 211.5l256-19.2l256 19.2v89l-254.6 20.7L0 300.5z"/><path fill="#eee" d="M0 178h512v33.5H0zm0 122.5h512V334H0z"/></g>`),
		g.Group(children),
	)
}