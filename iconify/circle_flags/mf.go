package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsMf0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsMf0)"><path fill="#eee" d="M0 128V0h512v128L299 512h-86z"/><circle cx="256" cy="213" r="57" fill="#ffda44"/><path fill="#eee" d="M185 213h142l-71 128Z"/><path fill="#d80027" d="M256 341L142 235h228z"/><path fill="#0052b4" d="m0 128l213 213v171H0Zm512 0L299 341v171h213z"/></g>`),
		g.Group(children),
	)
}