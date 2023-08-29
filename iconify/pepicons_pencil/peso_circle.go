package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PesoCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.5 5.5h-5v-1h5A4.5 4.5 0 0 1 18 9v1a4.5 4.5 0 0 1-4.5 4.5h-5v-1h5A3.5 3.5 0 0 0 17 10V9a3.5 3.5 0 0 0-3.5-3.5Z"/><path d="M8.5 4.5A.5.5 0 0 1 9 5v15.5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Z"/><path d="M5 7.936a.5.5 0 0 1 .5-.5h14a.5.5 0 0 1 0 1h-14a.5.5 0 0 1-.5-.5Zm0 3a.5.5 0 0 1 .5-.5h14a.5.5 0 0 1 0 1h-14a.5.5 0 0 1-.5-.5Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}