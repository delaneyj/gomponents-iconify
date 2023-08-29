package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsKw0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsKw0)"><path fill="#eee" d="M138.4 147L512 167v178l-373.6 20z"/><path fill="#6da544" d="m0 0l138.4 167H512V0z"/><path fill="#d80027" d="m0 512l138.4-167H512v167z"/><path fill="#333" d="M167 167L0 0v512l167-167z"/></g>`),
		g.Group(children),
	)
}