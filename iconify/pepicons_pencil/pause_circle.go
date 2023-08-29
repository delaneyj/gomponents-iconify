package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11 17V9a1 1 0 0 0-2 0v8a1 1 0 1 0 2 0ZM10 7a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V9a2 2 0 0 0-2-2Zm7 10V9a1 1 0 1 0-2 0v8a1 1 0 1 0 2 0ZM16 7a2 2 0 0 0-2 2v8a2 2 0 1 0 4 0V9a2 2 0 0 0-2-2Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}