package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmFrameOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.5 2.5a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h17a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1h-17Zm14 2h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1A.5.5 0 0 1 15 6V5a.5.5 0 0 1 .5-.5Zm1 3h-1a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1A.5.5 0 0 0 17 9V8a.5.5 0 0 0-.5-.5Zm-1 3h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5Zm1 3h-1a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5Zm-13-9h1A.5.5 0 0 1 5 5v1a.5.5 0 0 1-.5.5h-1A.5.5 0 0 1 3 6V5a.5.5 0 0 1 .5-.5Zm1 3h-1A.5.5 0 0 0 3 8v1a.5.5 0 0 0 .5.5h1A.5.5 0 0 0 5 9V8a.5.5 0 0 0-.5-.5Zm-1 3h1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-1A.5.5 0 0 1 3 12v-1a.5.5 0 0 1 .5-.5Zm1 3h-1a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1A.5.5 0 0 0 5 15v-1a.5.5 0 0 0-.5-.5Zm2 2v-11h7v11h-7Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}