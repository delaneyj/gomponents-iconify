package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func People(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4 8.25a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M4.05 11a1.5 1.5 0 0 0-1.5 1.5V14a1 1 0 0 1-2 0v-1.5a3.5 3.5 0 0 1 7 0V14a1 1 0 1 1-2 0v-1.5a1.5 1.5 0 0 0-1.5-1.5ZM16 8.25a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm0 2a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/><path d="M15.95 11a1.5 1.5 0 0 1 1.5 1.5V14a1 1 0 1 0 2 0v-1.5a3.5 3.5 0 0 0-7 0V14a1 1 0 1 0 2 0v-1.5a1.5 1.5 0 0 1 1.5-1.5Z"/><path d="M10.05 13.75a2.5 2.5 0 0 0-2.5 2.5v1.5a1 1 0 0 1-2 0v-1.5a4.5 4.5 0 0 1 9 0v1.5a1 1 0 1 1-2 0v-1.5a2.5 2.5 0 0 0-2.5-2.5Z"/><path d="M10 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g>`),
		g.Group(children),
	)
}