package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InGj(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsInGj0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsInGj0)"><path fill="#eee" d="M256 96L0 128v128l256 32l256-32V128L256 96zm0 256L0 384v128h512V384l-256-32z"/><path fill="#d80027" d="M0 0h512v128H0zm0 256h512v128H0z"/></g>`),
		g.Group(children),
	)
}