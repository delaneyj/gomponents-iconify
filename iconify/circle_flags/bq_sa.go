package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BqSa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsBqSa0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsBqSa0)"><path fill="#d80027" d="M0 0h512v256H0z"/><path fill="#0052b4" d="M0 256h512v256H0z"/><path fill="#eee" d="M0 256L256 0l256 256l-256 256z"/><path fill="#ffda44" d="m256 133.6l27.6 85H373L300.7 271l27.6 85l-72.3-52.5l-72.3 52.6l27.6-85l-72.3-52.6h89.4z"/></g>`),
		g.Group(children),
	)
}