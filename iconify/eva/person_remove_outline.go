package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonRemoveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPersonRemoveOutline0"><g id="evaPersonRemoveOutline1"><path id="evaPersonRemoveOutline2" fill="currentColor" d="M21 6h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2Zm-11 5a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm0 8a7 7 0 0 0-7 7a1 1 0 0 0 2 0a5 5 0 0 1 10 0a1 1 0 0 0 2 0a7 7 0 0 0-7-7Z"/></g></g>`),
		g.Group(children),
	)
}