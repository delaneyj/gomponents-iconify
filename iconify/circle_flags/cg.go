package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsCg0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsCg0)"><path fill="#ffda44" d="M384 0h128v128L352 352L128 512H0V384l160-224Z"/><path fill="#6da544" d="M0 384L384 0H0Z"/><path fill="#d80027" d="M512 128L128 512h384z"/></g>`),
		g.Group(children),
	)
}