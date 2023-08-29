package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsDc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsUsDc0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsUsDc0)"><path fill="#eee" d="M0 0h512v186l-64 48l64 48v57l-64 48l64 48v77H0v-77l64-48l-64-48v-57l64-48l-64-48Z"/><path fill="#d80027" d="M0 186h512v96H0zm0 153h512v96H0zm224-181l83-60H205l83 60l-32-98Zm118 0l83-60H323l83 60l-32-98Zm-236 0l83-60H87l83 60l-32-98Z"/></g>`),
		g.Group(children),
	)
}