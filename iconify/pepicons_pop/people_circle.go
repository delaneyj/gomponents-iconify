package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7 11.25a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M7.05 14a1.5 1.5 0 0 0-1.5 1.5V17a1 1 0 0 1-2 0v-1.5a3.5 3.5 0 0 1 7 0V17a1 1 0 1 1-2 0v-1.5a1.5 1.5 0 0 0-1.5-1.5ZM19 11.25a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm0 2a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/><path d="M18.95 14a1.5 1.5 0 0 1 1.5 1.5V17a1 1 0 1 0 2 0v-1.5a3.5 3.5 0 0 0-7 0V17a1 1 0 1 0 2 0v-1.5a1.5 1.5 0 0 1 1.5-1.5Z"/><path d="M13.05 16.75a2.5 2.5 0 0 0-2.5 2.5v1.5a1 1 0 0 1-2 0v-1.5a4.5 4.5 0 0 1 9 0v1.5a1 1 0 1 1-2 0v-1.5a2.5 2.5 0 0 0-2.5-2.5Z"/><path d="M13 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}