package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M2.5 18.5a.5.5 0 0 1 0-1h16a.5.5 0 0 1 0 1h-16Z"/><path d="M6.5 17.5a.5.5 0 0 1-1 0V4.808C5.5 3.272 6.602 2 8 2h5c1.398 0 2.5 1.272 2.5 2.808V17.5a.5.5 0 0 1-1 0V4.808C14.5 3.795 13.811 3 13 3H8c-.811 0-1.5.795-1.5 1.808V17.5Z"/><path d="M8.5 4.5h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1A.5.5 0 0 1 8 6V5a.5.5 0 0 1 .5-.5Zm3 0h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1A.5.5 0 0 1 11 6V5a.5.5 0 0 1 .5-.5Zm0 3h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1A.5.5 0 0 1 11 9V8a.5.5 0 0 1 .5-.5Zm0 3h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5Zm0 3h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5Zm-3-6h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1A.5.5 0 0 1 8 9V8a.5.5 0 0 1 .5-.5Zm0 6h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1A.5.5 0 0 1 8 15v-1a.5.5 0 0 1 .5-.5Zm0-3h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1A.5.5 0 0 1 8 12v-1a.5.5 0 0 1 .5-.5Z"/></g>`),
		g.Group(children),
	)
}